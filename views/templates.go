package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	HtmlTpl *template.Template
}

func ParseTheTemplate(filepath string) (Template, error) {
	parsedTemplate, parsingError := template.ParseFiles(filepath)
	if parsingError != nil {
		return Template{}, fmt.Errorf("parsing error on template: %w", parsingError)
	}

	return Template{
		HtmlTpl: parsedTemplate,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-type", "text/html; charset=utf-8")

	executeError := t.HtmlTpl.Execute(w, data)
	if executeError != nil {
		log.Printf("execute template error: %v", executeError.Error())
		http.Error(w, "there was error while executing the template", http.StatusInternalServerError)
		return
	}
}
