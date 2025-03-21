test:
	go test . -v

build:
	go build -o bin/$(APP_NAME) main.go

run:
	go run main.go

db-down:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost/point-of-sales-golang?sslmode=disable down

db-up:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost/point-of-sales-golang?sslmode=disable up

