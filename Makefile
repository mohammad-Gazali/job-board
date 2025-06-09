run: build
	@./bin/job-board

build-swagger:
	@swag init -g cmd/api/main.go

build:
	@go build -o bin/job-board cmd/api/main.go

sqlc:
	@sqlc generate -f db/sqlc.yaml

migrate:
	@goose up