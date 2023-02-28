hello:
	echo "Hello"

tidy:
	go mod tidy

build:
	go build -o bin/main main.go

swag:
	swag init --parseDependency  --parseInternal --parseDepth 1

run:
	go run main.go