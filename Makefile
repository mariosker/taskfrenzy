build:
	@go build -o bin/taskfrenzy cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/taskfrenzy

migration:
	@migrate create -ext psql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down