package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	// Servir les fichiers statiques du répertoire "static"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Home)
	http.HandleFunc("/login" , Login)
	fmt.Println("server started on port" + port)
	fmt.Println("http://localhost" + port)
	http.ListenAndServe(port, nil)
}
