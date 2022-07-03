# Usage

```bash
make setup

# firt start go metrics collector
make run-collector

# then start speedtest
make run-speedtest
```

# Grafana conf
Create a new dashboard on [grafana] and update the panel's field on the JSON Model with the data from grafana/dash.json file.

Default user and password are admin/admin.
Remeber to add a new Prometheus datasource as well. Set the URL as http://localhost:9090

[grafana]:http://localhost:3000