package main

import (
	"github.com/Post-and-Play/gears/infra"
	"github.com/joho/godotenv"
)

//v0.1

func main() {
	godotenv.Load(".env")
	infra.DatabaseConnect()
}
