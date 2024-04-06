build:
	@go build -o bin/ecommerce cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecommerce

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cm/migrate/main.go up

migrate-down:
	@go run cm/migrate/main.go down