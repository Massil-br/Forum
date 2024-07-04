package src

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type User struct {
	idUser   int
	username string
	email    string
	password string
}

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

	createCategoryTable:=`CREATE TABLE IF NOT EXISTS categories (
		"idCategory" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"idCategoryCreator" integer,
		"name" TEXT
	)`

	statement, err =db.Prepare(createCategoryTable)
	if err != nil{
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("Category Table created")

	createPostTable := `CREATE TABLE IF NOT EXISTS posts(
		"idPost" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"categoryID" integer,
		"idPostCreator" integer,
		"postTitle" TEXT,
		"postContent" TEXT,
		"likes" integer,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
	)`

	statement , err = db.Prepare(createPostTable)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("Post Table created")

	createCommentTable := `CREATE TABLE IF NOT EXISTS comments(
		"idComment" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"idCommentCreator" integer,
		"idPostOfComment" integer,
		"commentContent" TEXT,
		"likes" integer,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
	)`

	statement , err = db.Prepare(createCommentTable)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("Comments Table created")

	createFavoriteCategoryTable := `CREATE TABLE IF NOT EXISTS favoriteCategories(
		"idCategory" integer,
		"idUser" integer
	)`

	statement, err = db.Prepare(createFavoriteCategoryTable)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("Favorite Categories Table created")

	createLikesTable := `CREATE TABLE IF NOT EXISTS likes(
		"idLike" integer NOT NULL PRIMARY KEY AUTOINCREMENT, 
		"idUser" integer,
		"idPost" integer
	)`

	statement, err = db.Prepare(createLikesTable)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("Likes Table created")

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

func ShowDatabase() {
	rows, _ := db.Query("SELECT idUser, username, email, password FROM users ")
	thisUser := User{}
	for rows.Next() {
		rows.Scan(&thisUser.idUser, &thisUser.username, &thisUser.email, &thisUser.password)
		fmt.Println(thisUser)
	}
}

func CheckIfUserExist(username string, email string) (bool, int) {
	boolean := false
	var userId int

	rows, _ := db.Query("SELECT idUser, username, email, password FROM users")
	thisUser := User{}
	for rows.Next() {
		rows.Scan(&thisUser.idUser, &thisUser.username, &thisUser.email, &thisUser.password)
		if thisUser.username == username || thisUser.email == email {
			boolean = true
			userId = thisUser.idUser
		}
	}
	return boolean, userId
}

func GetUserByID(ID int) User {
	rows, _ := db.Query("SELECT idUser, username, email, password FROM users")
	userToReturn := User{}
	user := User{}
	for rows.Next() {
		rows.Scan(&user.idUser, &user.username, &user.email, &user.password)
		if user.idUser == ID {
			userToReturn = user
		}
	}
	return userToReturn
}

func (user *User) GetUsername() string {
	return user.username
}

func (user *User) GetID() int {
	return user.idUser
}

func (user *User) GetEmail() string{
	return user.email
}
