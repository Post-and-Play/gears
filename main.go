package main

import (
	"github.com/Post-and-Play/gears/infra"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	infra.DatabaseConnect()
}
