package main

import (
	"html/template"
	"net/http"

)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)

}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func Login(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w , "login")
}

func Register(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register")
}

func Categories(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "categories")
}