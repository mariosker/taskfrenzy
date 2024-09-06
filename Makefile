build:
	@go build -o bin/taskfrenzy cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/taskfrenzy
