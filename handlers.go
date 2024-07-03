package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/Massil-br/Forum.git/src"
)

var sessions = map[string]int{}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func Home(w http.ResponseWriter, r *http.Request) {
	user := getUserFromSession(r)
	renderTemplate(w, "home", user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("loginUsername")
		password := r.FormValue("loginPassword")

		exist, ID := src.CheckIfUserExist(username, username)
		if !exist {
			fmt.Println("Incorrect Username or password")
		}
		User := src.GetUserByID(ID)

		match := User.CheckPassword(password)
	

		if match {
			fmt.Println("Connected as ", User.GetUsername())
			setUserSession(w, &User)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			fmt.Println("Incorrect Username or password")
		}
	}
	user := getUserFromSession(r)
	renderTemplate(w, "login", user)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm_password")

		exist, _ := src.CheckIfUserExist(username, email)

		if exist {
			fmt.Println("username or email already taken ")
			renderTemplate(w, "register", nil)
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
			return
		} else {
			fmt.Println("password and confirmpassword are not the same")
		}
	}
	renderTemplate(w, "register", nil)
}

func Categories(w http.ResponseWriter, r *http.Request) {
	user := getUserFromSession(r)
	renderTemplate(w, "categories", user)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	user := getUserFromSession(r)
	renderTemplate(w, "create-category", user)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	user := getUserFromSession(r)
	renderTemplate(w, "create-post", user)
}

func getUserFromSession(r *http.Request) *src.User {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return nil
	}
	userID, ok := sessions[cookie.Value]
	if !ok {
		return nil
	}
	user := src.GetUserByID(userID)
	return &user
}

func setUserSession(w http.ResponseWriter, user *src.User) {
	sessionToken := fmt.Sprintf("%d", time.Now().UnixNano())
	sessions[sessionToken] = user.GetID()
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(24 * time.Hour),
	})
}