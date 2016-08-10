package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Loading up env variables")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Listening on http://localhost:3000")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
