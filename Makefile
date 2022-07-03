setup:
	@go mod tidy

	@docker-compose down -v
	@docker-compose up -d

run-collector:
	@go run main.go

run-speedtest:
	./scripts/output-speedtest-json.sh

.PHONY: .setup .run-speedtest