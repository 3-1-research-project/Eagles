package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"minitwit.com/db"
	"minitwit.com/model"
	"minitwit.com/sim"
)

func Follow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	sim.UpdateLatest(r)

	var rv model.FollowData
	err := json.NewDecoder(r.Body).Decode(&rv)
	if err != nil {
		fmt.Println("Error in decoding the JSON", err)
	}
	is_auth := db.Is_authenticated(w, r)
	if !is_auth {
		return
	}
	user_id, _ := db.Get_user_id(username)
	if db.IsNil(user_id) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	no_flws := no_followees(r, 100)

	if r.Method == "POST" && rv.Follow != "" {

		follow_username := rv.Follow
		follow_user_id, _ := db.Get_user_id(follow_username)

		if db.IsNil(follow_user_id) {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		query := `INSERT INTO follower (who_id, whom_id) VALUES (?, ?)`
		sqlite_db, _ := db.Connect_db()
		defer sqlite_db.Close()
		_, err := sqlite_db.Exec(query, user_id, follow_user_id)

		if err != nil {
			fmt.Println("Error querying the database")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(http.StatusOK)

	} else if r.Method == "POST" && rv.Unfollow != "" {

		unfollow_username := rv.Unfollow
		unfollow_user_id, err := db.Get_user_id(unfollow_username)

		if err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		query := `DELETE FROM follower WHERE who_id=? and WHOM_id=?`
		sqlite_db, _ := db.Connect_db()
		defer sqlite_db.Close()
		_, err = sqlite_db.Exec(query, user_id, unfollow_user_id)

		json.NewEncoder(w).Encode(http.StatusOK)

	} else if r.Method == "GET" {
		followees := db.GetFollowees([]any{user_id, no_flws}, false)

		json.NewEncoder(w).Encode(struct {
			Follows []string `json:"follows"`
		}{
			Follows: followees,
		})
	}
}

func no_followees(r *http.Request, defaultValue int) int {
	value := r.URL.Query().Get("no")
	if value != "" {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return intValue
		}
	}
	return defaultValue
}
