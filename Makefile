test:
	go test . -v

build:
	go build -o bin/$(APP_NAME) main.go

run:
	air

create-migrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

db-down:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost/point-of-sales-golang?sslmode=disable down

db-up:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost/point-of-sales-golang?sslmode=disable up

db-test-down:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost/point-of-sales-golang-test?sslmode=disable down

db-test-up:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost/point-of-sales-golang-test?sslmode=disable up

install:
	go mod download
