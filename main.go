package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	println("Hello, World!$PORT must be set")	
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	println("PORT is set to " + port)
}