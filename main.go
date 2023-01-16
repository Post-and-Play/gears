package main

import (
	"github.com/Johnman67112/gears/infra"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	infra.DatabaseConnect()
}
