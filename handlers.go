package main

import (
	"html/template"
	"net/http"
)

func (fe *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{}); err != nil {
		panic(err.Error())
	}
}

var (
	templates = template.Must(template.New("").
		ParseGlob("templates/*.html"))
)
