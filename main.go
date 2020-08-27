package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("docker/database.env")

	a := App{}
	a.Initialize(os.Getenv("DB_CONNECTION_URL"))

	a.Run(":8010")
}
