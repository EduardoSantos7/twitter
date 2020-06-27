package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Client connected to DB*/
var MongoCN = DBConnection()
var clientOptions = options.Client().ApplyURI("mongodb+srv://eduardo:05120714@cluster0-kle6z.azure.mongodb.net/twitter?retryWrites=true&w=majority")

/* DBConnection start the connection with DB */
func DBConnection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connected to DB")

	return client
}

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}
	return 1
}
