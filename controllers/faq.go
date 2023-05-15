package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Faq(w http.ResponseWriter, r *http.Request) {

	// build the data
	type Questions []struct {
		Question string
		Answer   string
	}

	questions := Questions{
		{
			Question: "what is your name",
			Answer:   "my name is khan",
		},
		{
			Question: "How old are you",
			Answer:   "iam 21",
		},
	}

	faqHtml, faqhtmlError := template.ParseFiles("views/components/layouts.gohtml", "views/pages/faq.gohtml")

	if faqhtmlError != nil {
		log.Printf("Faq html rendering error: %v", faqhtmlError.Error())
		return
	}

	faqHtml.Execute(w, questions)

}
