package main

import (
	"fmt"
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
		confirmPassword := r.FormValue("confirm_password")

		exist := src.CheckIfUserExist(username, email)

		if exist {
			fmt.Println("username or email already taken ")
			renderTemplate(w, "register") // Added to prevent inserting new data
			return
		}

		hashedPassword, err := src.HashPassword(password)
		if err != nil {
			fmt.Println("error while hashing password")
		}

		match := src.CheckPasswordHash(confirmPassword, hashedPassword)

		if match {
			src.InsertUser(username, email, hashedPassword)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			src.ShowDatabase()
			return
		} else {
			fmt.Println("password and confirmpassword are not the same")
		}
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
