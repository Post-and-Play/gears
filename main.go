package main

import (
	"fmt"
	"os"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/ui"
	"github.com/joho/godotenv"
)

//v0.1

func main() {
	godotenv.Load(".env")
	fmt.Println(os.Getenv("DB_HOST"))
	infra.DatabaseConnect()
	ui.Router()
}
