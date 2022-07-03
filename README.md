# Usage

```bash
make setup

# firt start go metrics collector
make run-collector

# then start speedtest
make run-speedtest
```

# Grafana conf
Default user and password are admin/admin.
Remeber to add a new Prometheus datasource as well. Set the URL as http://localhost:9090

Create a new dashboard on [grafana] and update the panel's field on the JSON Model with the data from the grafana/dash.json file. The uids for the datasources (on the dash.json file) have to be update to match the data source you created before.

[grafana]:http://localhost:3000