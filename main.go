package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("docker/database.env")

	a := App{}
	a.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	a.Run(":8010")
}
