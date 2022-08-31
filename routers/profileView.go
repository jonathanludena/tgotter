package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanludena/tgotter/bd"
)

/* Extract data profile */
func ProfileView(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Id must be required in params", http.StatusBadRequest)
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Error searching profile "+err.Error(), 400)
		return
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
