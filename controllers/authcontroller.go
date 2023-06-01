package controllers

import (
	"fmt"
	"gocourse/models"
	"html/template"
	"net/http"
)

type Users struct {
	UserService *models.UserService
}

func SignUpGet(w http.ResponseWriter, r *http.Request) {
	// parse the file

	signUpHtml, htmlError := template.ParseFiles("views/components/layouts.gohtml", "views/pages/auth/signup.gohtml")
	if htmlError != nil {
		fmt.Errorf("error while parsing the signup page: %v", htmlError)
		return
	}

	singupExecuteError := signUpHtml.Execute(w, nil)

	if singupExecuteError != nil {
		fmt.Errorf("error while parsing the signup page: %v", singupExecuteError)
		return
	}
}

func (u Users) SignUpPost(w http.ResponseWriter, r *http.Request) {

	requestEmail := r.FormValue("email")
	requestPassword := r.FormValue("password")

	user, err := u.UserService.Create(requestEmail, requestPassword)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "User Created :%+v", user)

}
