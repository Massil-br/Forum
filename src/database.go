package src

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"idUser" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"username" TEXT,
		"email" TEXT,
		"password" TEXT		
	  );`

	statement, err := db.Prepare(createUserTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("User table created")
}

func InsertUser(username, email, password string) {
	insertUserSQL := `INSERT INTO users(username, email, password) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL) // Utiliser src.db
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(username, email, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted user successfully")
}
