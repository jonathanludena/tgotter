package routers

import (
	"errors"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jonathanludena/tgotter/bd"
	"github.com/jonathanludena/tgotter/models"
)

/* Email used in all API endpoints */
var Email string

/* user ID returned of model, it will use in all API endpoints */
var IDUser string

/* Process Token extract data token */
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	SECRET := os.Getenv("SECRET")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer ")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if err == nil {
		_, finded, _ := bd.CheckUserExists(claims.Email)
		if finded {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, finded, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
