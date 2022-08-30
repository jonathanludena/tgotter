package bd

import (
	"github.com/jonathanludena/tgotter/models"

	"golang.org/x/crypto/bcrypt"
)

/* Function Login and It will try signin with data user in DB */
func TryLogin(email string, password string) (models.User, bool) {
	usu, finded, _ := CheckUserExists(email)
	if !finded {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}

	return usu, true
}
