package main

import (
	"fmt"
	"net/http"

	"github.com/Massil-br/Forum.git/src"
)

const port = ":8080"

func main() {
	src.InitDB() // Initialiser la base de données
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/categories", Categories)
	http.HandleFunc("/create-category", CreateCategory)
	http.HandleFunc("/create-post", CreatePost)
	http.HandleFunc("/profile", Profile)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/postlist/", PostList) 
	http.HandleFunc("/posts/", PostContent)

	fmt.Println("server started on port" + port)
	fmt.Println("http://localhost" + port)
	http.ListenAndServe(port, nil)
}