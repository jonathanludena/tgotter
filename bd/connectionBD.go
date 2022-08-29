package bd

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Object Connection to DB */
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI(os.Getenv("URI_DB"))

/* Connect function to DB */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection Success to DB")
	return client
}

/* Check function connection with DB */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
