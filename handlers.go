package main

import (
	"html/template"
	"net/http"

	"github.com/Massil-br/Forum.git/src"
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
	renderTemplate(w, "login")
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		src.InsertUser(username, email, password)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	renderTemplate(w, "register")
}

func Categories(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "categories")
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "create-category")
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "create-post")
}
