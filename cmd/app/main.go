package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    handler := http.NewServeMux()
    handler.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    log.Fatal(http.ListenAndServe(":8080", handler))
}
