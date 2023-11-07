package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nobilik/bgwork/RSSReaderService/handlers"
)

func init() {
	if os.Getenv("API_PORT") != "" {
		return
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func main() {
	mux := handlers.SetupRoutes()
	fmt.Printf("Server listening on port %s\n", os.Getenv("API_PORT"))
	http.ListenAndServe(":"+os.Getenv("API_PORT"), mux)
}
