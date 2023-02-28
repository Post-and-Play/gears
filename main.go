package main

import (
	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/ui"
	"github.com/joho/godotenv"
)

//v0.3

func main() {
	godotenv.Load(".env")
	infra.DatabaseConnect()
	ui.RunServer()
}
