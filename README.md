# Dependencies
GO: https://go.dev/doc/install

Docker: https://docs.docker.com/engine/install/

Docker Compose: https://docs.docker.com/compose/install/

Speedtest: https://www.speedtest.net/apps/cli

# Usage

```bash
make setup

# firt start go metrics collector
make run-collector

# then start speedtest
make run-speedtest
```

# Grafana
http://localhost:3000

Default user and password are admin/admin.
You should see a datasource named Speedtest.

# Prometheus
http://localhost:9090

# Metrics endpoint
http://localhost:3001/metrics
