package main

import (
	"fmt"
	"gocourse/controllers"
	"gocourse/models"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")

	contacthtml, contact_html_error := template.ParseFiles("views/components/layouts.gohtml", "views/pages/contact.gohtml")
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

	homehtml, homehtmlError := template.ParseFiles("views/components/layouts.gohtml", "views/pages/home.gohtml")

	if homehtmlError != nil {
		log.Printf("error while parsing home.gohtml: %v", homehtmlError)
		http.Error(w, "error while parsing home.gohtml", http.StatusInternalServerError)
		return
	}

	homehtml.Execute(w, nil)

}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", controllers.Faq)

	// authentication
	r.Get("/signup", controllers.SignUpGet)
	r.Get("/login", controllers.LgoinGet)

	config := models.DefaultPostgressConfig()
	db, err := models.Open(config)

	if err != nil {
		fmt.Errorf("error while opening the DB : %w", err)
	}

	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	userController := controllers.Users{
		UserService: &userService,
	}

	r.Post("/user-create", userController.SignUpPost)
	r.Post("/login-post", userController.LoginPost)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})

	fmt.Println("server is running on :3000")
	http.ListenAndServe(":3000", r)
}
