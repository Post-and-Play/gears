package main

import (
	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/ui"
	"github.com/joho/godotenv"
)

//v0.4

func main() {
	godotenv.Load(".env")
	infra.DatabaseConnect()
	ui.RunServer()
}
