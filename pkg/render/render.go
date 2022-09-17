package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// Rendertemplate creates a new template from the passed files
func RenderTemplateTest(w http.ResponseWriter, tmplName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmplName, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	var template *template.Template

	//check to see if the template is already in the cache
	_, inMap := templateCache[templateName]

	if !inMap {
		//need to create a new template
		addTemplateToCache(templateName)
	} else {
		//get the template from the cache
		template = templateCache[templateName]

	}

	err := template.Execute(w, nil)
	println(err)
}

func addTemplateToCache(templateName string) error {
	sliceOfTemplates := []string{
		fmt.Sprintf("./templates/%s", templateName),
		"./templates/base.layout.tmpl",
	}

	//parse the template
	parsedTemplate, err := template.ParseFiles(sliceOfTemplates...)

	if err != nil {
		//add the new template to the cache
		templateCache[templateName] = parsedTemplate
		return nil
	}
	//else return the error
	return err
}
