package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+port, router))
}
