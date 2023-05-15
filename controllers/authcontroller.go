package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

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

func SignUpPost(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "email : ", r.FormValue("email"))
	fmt.Fprintln(w, "password : ", r.FormValue("password"))

}
