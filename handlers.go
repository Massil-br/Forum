package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Massil-br/Forum.git/class"
	"github.com/Massil-br/Forum.git/src"
	"github.com/gofrs/uuid"
)

var sessions = map[string]int{}

type Server struct {
	User       *class.User
	Categories []class.Category
	Posts      []class.Post 
	Post *class.Post

}

var server = Server{}

func Home(w http.ResponseWriter, r *http.Request) {
	user := getUserFromSession(r)
	if user != nil {
		server.User = user
	} else {
		server.User = nil
	}
	renderTemplate(w, "home", server)
}

func Login(w http.ResponseWriter, r *http.Request) {
	login(w, r)
	user := getUserFromSession(r)
	if user != nil {
		server.User = user
	} else {
		server.User = nil
	}
	renderTemplate(w, "login", server)
}

func Register(w http.ResponseWriter, r *http.Request) {
	register(w, r)
	user := getUserFromSession(r)
	if user != nil {
		server.User = user
	} else {
		server.User = nil
	}
	renderTemplate(w, "register", server)
}

func Categories(w http.ResponseWriter, r *http.Request) {

	user := getUserFromSession(r)
	if user != nil {
		server.User = user
	} else {
		server.User = nil
	}
	server.Categories = src.GetCategories()

	renderTemplate(w, "categories", server)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	user := getUserFromSession(r)
	if user != nil {
		server.User = user
	} else {
		server.User = nil
	}
	renderTemplate(w, "profile", server)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	createCategory(w, r)
	user := getUserFromSession(r)
	if user != nil {
		server.User = user
	} else {
		server.User = nil
	}
	renderTemplate(w, "create-category", server)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	createPost(w, r)
	user := getUserFromSession(r)
	if user != nil {
		server.User = user
	} else {
		server.User = nil
	}
	server.Categories = src.GetCategories()

	renderTemplate(w, "create-post", server)
}

func PostList(w http.ResponseWriter, r *http.Request) {
	user := getUserFromSession(r)
	if user != nil {
		server.User = user
	} else {
		server.User = nil
	}

	// Récupérer l'ID de la catégorie à partir de l'URL
	categoryIDStr := strings.TrimPrefix(r.URL.Path, "/postlist/")
	categoryIDStr, err := url.PathUnescape(categoryIDStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		fmt.Println("Error unescaping category ID:", err)
		return
	}

	// Convertir l'ID de la catégorie en entier
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		fmt.Println("Error converting category ID to integer:", err)
		return
	}

	
	server.Posts, err = src.GetPostsByID(categoryID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		return
	}

	renderTemplate(w, "post-list", server)
}


func PostContent(w http.ResponseWriter, r *http.Request){
	user := getUserFromSession(r)
	if user != nil {
		server.User = user
	}else {
		server.User = nil
	}

	postIDStr := strings.TrimPrefix(r.URL.Path, "/posts/")
	postIDStr, err := url.PathUnescape(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		fmt.Println("Error unescaping post ID:", err)
		return
	}
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		fmt.Println("Error converting post ID to integer:", err)
		return
	}

	for _, post := range server.Posts {
		if post.GetID() == postID {
			server.Post = &post
		}
	}
	fmt.Println("Post to display:", server.Post) // Ajout de log

	renderTemplate(w, "post", server)
}

/* from here to the bottom
the code is not for loading
pages but only funcs that
are used in the handlers */

func getUserFromSession(r *http.Request) *class.User {
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

func setUserSession(w http.ResponseWriter, user *class.User) {
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

func register(w http.ResponseWriter, r *http.Request) {
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
}

func login(w http.ResponseWriter, r *http.Request) {
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
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		src.InsertCategory(name, server.User.GetID())
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		return
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		categoryID, err := strconv.Atoi(r.FormValue("category"))
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			log.Printf("Error converting category ID: %v", err)
			return
		}
		user := getUserFromSession(r)
		if user != nil {
			server.User = user
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		err = src.InsertPost(title, content, server.User.GetID(), categoryID)
		if err != nil {
			http.Error(w, "Error inserting post", http.StatusInternalServerError)
			log.Printf("Error inserting post: %v", err)
			return
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		return
	}
}