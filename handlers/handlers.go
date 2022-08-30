package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/jonathanludena/tgotter/middlew"
	"github.com/jonathanludena/tgotter/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handler API and Listen PORT */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT != "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
