package handlers

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static
var staticFS embed.FS

func (s Server) serveStatic() http.HandlerFunc {

    fileSys, err := fs.Sub(staticFS, "static")

    if err != nil {
        panic(err)
    }

    server := http.FileServer(http.FS(fileSys))

    return func(w http.ResponseWriter, r *http.Request) {
        server.ServeHTTP(w, r)
    }
}

