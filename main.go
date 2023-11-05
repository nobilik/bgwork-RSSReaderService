package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("PORT") != "" {
		return
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func main() {
	server()
}
