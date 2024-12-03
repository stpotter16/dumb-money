package handlers

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates
var templateFS embed.FS

var baseTemplates = []string{
	"templates/layouts/base.html",
}

func (s Server) indexGet() http.HandlerFunc {
	t := template.Must(template.New("base.html").ParseFS(templateFS, append(baseTemplates, "templates/pages/index.html")...))
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}
