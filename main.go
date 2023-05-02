package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")

	contacthtml, contact_html_error := template.ParseFiles("templates/contact.gohtml")
	if contact_html_error != nil {
		log.Printf("Error while parsing the contacnt html: %v\n", contact_html_error.Error())
		http.Error(w, "error on contact html file", http.StatusInternalServerError)
	}

	executeError := contacthtml.Execute(w, nil)

	if executeError != nil {
		log.Printf("Error while executing the contacnt html: %v\n", executeError.Error())
		http.Error(w, "error on executing html file", http.StatusInternalServerError)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")

	// filepath := filepath.Join("templates", "home.gohtml") // another way of parsing files

	homehtml, homehtmlError := template.ParseFiles("templates/home.gohtml")
	if homehtmlError != nil {
		log.Printf("error while parsing home.gohtml: %v", homehtmlError)
		http.Error(w, "error while parsing home.gohtml", http.StatusInternalServerError)
		return
	}

	templateError := homehtml.Execute(w, nil)

	if templateError != nil {
		panic(templateError.Error())
	}

}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})

	fmt.Println("server is running on :3000")
	http.ListenAndServe(":3000", r)
}
