package handlers

import (
	"fmt"
	"net/http"
)

func (s Server) indexGet() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    }
}
