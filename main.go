package main

import (
	"os"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/ui"
	"github.com/joho/godotenv"
)

//v0.6

func main() {
	godotenv.Load(".env")

	env := os.Getenv("APP_ENV")
	if env == "PROD" {
		infra.ProdDatabaseConnect()
	} else {
		infra.DevDatabaseConnect()
	}

	ui.RunServer()
}
