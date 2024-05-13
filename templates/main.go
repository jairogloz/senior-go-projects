package main

import (
	"html/template"
	"net/http"
)

type PageVariables struct {
	Name    string
	Message string
}

func main() {
	http.HandleFunc("/", HomePage)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	HomePageVars := PageVariables{
		Name:    "Senior Go Español!!!",
		Message: "Ya tenemos comunidad en Slack!",
	}

	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, HomePageVars)
	if err != nil {
		panic(err)
	}
}
