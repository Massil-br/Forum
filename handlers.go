package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/Massil-br/Forum.git/src"
	"github.com/gofrs/uuid" // Added uuid package
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

func Profile(w http.ResponseWriter, r *http.Request){
	user := getUserFromSession(r)
	renderTemplate(w, "profile", user)
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
	// Generate a UUID for the session token
	sessionToken, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Error generating session token", http.StatusInternalServerError)
		return
	}
	sessions[sessionToken.String()] = user.GetID()
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken.String(),
		Expires: time.Now().Add(10 * time.Minute),
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_token")
    if err != nil {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Supprimer la session
    delete(sessions, cookie.Value)

    // Expirer le cookie
    http.SetCookie(w, &http.Cookie{
        Name:    "session_token",
        Value:   "",
        Expires: time.Now().Add(-1 * time.Hour),
    })

    http.Redirect(w, r, "/login", http.StatusSeeOther)
}