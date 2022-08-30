package bd

import "golang.org/x/crypto/bcrypt"

/* Function Encrypt Password when save user to DB */
func EncryptPass(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)

	return string(bytes), err
}
