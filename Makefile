run: build
	@./bin/job-board

build:
	@go build -o bin/job-board cmd/api/main.go