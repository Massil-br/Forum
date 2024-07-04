package class

import "golang.org/x/crypto/bcrypt"

type User struct {
	idUser   int
	username string
	email    string
	password string
}

func (user *User) GetUsername() string {
	return user.username
}

func (user *User) GetID() int {
	return user.idUser
}

func (user *User) GetEmail() string {
	return user.email
}
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))
	return err == nil
}


func (user *User) GetIDAdress() *int {
	return &user.idUser
}

func (user *User) GetUsernameAdress() *string {
	return &user.username
}

func (user *User) GetEmailAdress() *string {
	return &user.email
}

func (user *User) GetPasswordAdress() *string {
	return &user.password
}
