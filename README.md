# Energy Measurement
To run this project, you first need to start a Postgres database (see https://github.com/3-1-research-project/EnergyMeasurementMVP) and then either use the VS code task `run-minitwit` or, in two separate terminals, run.

Backend terminal
```bash
source set-environment-variables.sh
cd minitwit-api/
go run minitwit-api.go
```

Frontend terminal
```bash
source set-environment-variables.sh
cd minitwit-web-app/
go run minitwit.go
```

# Welcome to the new MiniTwit app!

## Apps live here:
### Main app: 
http://188.166.201.66:15000

### API service: 
http://188.166.201.66:15001

### Grafana
http://188.166.201.66:3000/

### Prometheus
http://188.166.201.66:15001/metrics

## Folders
* **docs** documentation
* **legacy_tests** Python tests as provided from the legacy setup.
* **minitwit-web-app** The current version of MiniTwit written in Go
* **tests** The modernised tests written in Go

# Ways of working
[docs/working-process.md](docs/working-process.md)

# How to: 
[docs/how-to.md](docs/how-to.md)

# Status page
[LIVE APP STATUS](http://206.81.24.116/status.html)

*Authors:*
Daniel Grønbjerg
Roman Zvoda
Viktoria Hopeberg
Trond Rossing
Rasmus Balder Nordbjærg
Jan Lishak
