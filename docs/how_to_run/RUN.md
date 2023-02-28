# Config

## Setup

* Run go mod tidy

```  shell
    make tidy / go mod tidy
```

* Copy .env.example to .env file

* Docker-Compose for database

```  shell
    make compose / docker-compose up -d
```

* Running

```  shell
    make run / go run main.go
```

## Swag Update

```  shell
    make swag / swag init --parseDependency  --parseInternal --parseDepth 1
```

## Build

```  shell
    make build / go build -o bin/main main.go
```
