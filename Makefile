test:
	go test . -v

build:
	go build -o bin/$(APP_NAME) main.go

run:
	go run main.go