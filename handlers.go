package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

// Rendertemplate creates a new template from the passed files
func renderTemplate(w http.ResponseWriter, tmplName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmplName)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
