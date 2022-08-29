package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jonathanludena/tgotter/bd"
	"github.com/jonathanludena/tgotter/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Server without connection to DB")
		return
	}

	handlers.Handlers()
}
