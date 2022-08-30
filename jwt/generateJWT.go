package jwt

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jonathanludena/tgotter/models"
)

/* Generate Json Web Token */
func GenerateJWT(t models.User) (string, error) {
	SECRET := os.Getenv("SECRET")
	mySecret := []byte(SECRET)
	payload := jwt.MapClaims{
		"email":    t.Email,
		"name":     t.Name,
		"lastname": t.Lastname,
		"birthday": t.Birthday,
		"bio":      t.Bio,
		"location": t.Location,
		"web":      t.Siteweb,
		"_id":      t.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(mySecret)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
