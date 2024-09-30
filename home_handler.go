package main

import (
	"html/template"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Home() http.HandlerFunc {
	t := template.Must(template.ParseFS(viewsFS, "views/layout.html", "views/home.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}
