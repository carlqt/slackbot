package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	router := NewRouter()

	log.Println("listening to port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
