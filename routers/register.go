package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanludena/tgotter/bd"
	"github.com/jonathanludena/tgotter/models"
)

/* Function create a new user in DB */
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error data received "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must have min 6 characters", 400)
		return
	}

	_, finded, _ := bd.CheckUserExists(t.Email)
	if finded {
		http.Error(w, "User exists with this email", 400)
		return
	}

	_, status, err := bd.InsertUser(t)
	if err != nil {
		http.Error(w, "Error in operation => InsertUser", 400)
		return
	}

	if !status {
		http.Error(w, "Something was wrong, error in operation => InsertUser", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
