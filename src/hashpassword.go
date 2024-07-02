package src

import "golang.org/x/crypto/bcrypt"



func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword( []byte (password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err:=bcrypt.CompareHashAndPassword([]byte (hash), []byte (password))
	return err == nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte (user.password), []byte (password))
	return err == nil
}