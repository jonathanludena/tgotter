package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jonathanludena/tgotter/bd"
	"github.com/jonathanludena/tgotter/jwt"
	"github.com/jonathanludena/tgotter/models"
)

/* Function Login */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid User or Password "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	doc, exists := bd.TryLogin(t.Email, t.Password)
	if !exists {
		http.Error(w, "Invalid User or Password", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(doc)
	if err != nil {
		http.Error(w, "Something was wrong, please try again later "+err.Error(), 400)
		return
	}

	resp := models.RespLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
