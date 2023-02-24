hello:
	echo "Hello"

tidy:
	go mod tidy

build:
	go build -o bin/main main.go

run:
	go run main.go