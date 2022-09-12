package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// Rendertemplate creates a new template from the passed files
func RenderTemplate(w http.ResponseWriter, tmplName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmplName)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
