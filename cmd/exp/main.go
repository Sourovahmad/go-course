package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {
	html, templateError := template.ParseFiles("index.gohtml")

	if templateError != nil {
		panic(templateError.Error())
	}

	user := User{
		Name: "khan",
	}

	executeError := html.Execute(os.Stdout, user)

	if executeError != nil {
		panic(executeError.Error())
	}
}
