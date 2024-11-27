package main

import (
	"log"
	"net/http"

	"github.com/stpotter16/dumb-money/handlers"
)

func main() {
	server := handlers.New()
	log.Fatal(http.ListenAndServe(":8080", server.Handler()))
}
