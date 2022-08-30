package middlew

import (
	"net/http"

	"github.com/jonathanludena/tgotter/bd"
)

/* Check connection to DB and return next handlerFunc of ServerHttp */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Connection to DB losted", 500)
		}

		next.ServeHTTP(w, r)
	}
}
