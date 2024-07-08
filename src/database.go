package src

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Massil-br/Forum.git/class"
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

	createCategoryTable := `CREATE TABLE IF NOT EXISTS categories (
		"idCategory" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"idCategoryCreator" integer,
		"name" TEXT
	)`

	statement, err = db.Prepare(createCategoryTable)
	if err != nil {
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

	statement, err = db.Prepare(createPostTable)
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

	statement, err = db.Prepare(createCommentTable)
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
	thisUser := class.User{}
	for rows.Next() {
		rows.Scan(thisUser.GetIDAdress(), thisUser.GetUsernameAdress(), thisUser.GetEmailAdress(), thisUser.GetPasswordAdress())
		fmt.Println(thisUser)
	}
}

func CheckIfUserExist(username string, email string) (bool, int) {
	boolean := false
	var userId int

	rows, _ := db.Query("SELECT idUser, username, email, password FROM users")
	thisUser := class.User{}
	for rows.Next() {
		rows.Scan(thisUser.GetIDAdress(), thisUser.GetUsernameAdress(), thisUser.GetEmailAdress(), thisUser.GetPasswordAdress())
		if thisUser.GetUsername() == username || thisUser.GetEmail() == email {
			boolean = true
			userId = thisUser.GetID()
		}
	}
	return boolean, userId
}

func GetUserByID(ID int) class.User {
	rows, _ := db.Query("SELECT idUser, username, email, password FROM users")
	userToReturn := class.User{}
	user := class.User{}
	for rows.Next() {
		rows.Scan(user.GetIDAdress(), user.GetUsernameAdress(), user.GetEmailAdress(), user.GetPasswordAdress())
		if user.GetID() == ID {
			userToReturn = user
		}
	}
	return userToReturn
}

func InsertCategory(name string, userID int) {
	insertCategorySQL := `INSERT INTO categories(idCategoryCreator, name) VALUES (?, ?)`
	statement, err := db.Prepare(insertCategorySQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(userID, name)
}

func InsertPost(title string, content string, userID int, categoryID int) {
	insertPostSQL := `INSERT INTO posts(idPostCreator, categoryID, postTitle, postContent) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertPostSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(userID, categoryID, title, content)
}

func GetCategories() []class.Category {
	rows, _ := db.Query("SELECT idCategory, idCategoryCreator, name FROM categories")
	category := class.Category{}
	Categories := []class.Category{}
	for rows.Next() {
		rows.Scan(category.GetIDAdress(), category.GetCategoryCreatorIDAdress(), category.GetNameAdress())
		Categories = append(Categories, category)
	}
	fmt.Println(Categories)
	return Categories
}

func GetPostsByID(categoryID int) ([]class.Post, error) {
	var posts []class.Post

	rows, err := db.Query("SELECT idPost, categoryID, idPostCreator, postTitle, postContent, likes FROM posts WHERE categoryID = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post class.Post
		if err := rows.Scan(post.GetIDAdress(), post.GetIDCategoryAdress(), post.GetIDPostCreatorAdress(), post.GetPostTitleAdress(), post.GetPostContentAdress(), post.GetPostLikesAdress()); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, errors.New("no posts found for this category")
	}

	return posts, nil
}