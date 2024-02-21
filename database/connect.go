package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoURL = "YOUR MONGO URL HERE"
var ClientMongo = ConectarDB()
var MongBD = "persondb"
var CategoryColection = ClientMongo.Database(MongBD).Collection("categories")
var ProductColection = ClientMongo.Database(MongBD).Collection("products")
var ProductPhotoCollection = ClientMongo.Database(MongBD).Collection("photo_product")
var UsersCollection = ClientMongo.Database(MongBD).Collection("users")
var clientOptions = options.Client().ApplyURI(MongoURL + MongBD)

func ConectarDB() *mongo.Client {
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
	log.Println("Connected Successfully to MONGO")
	return client
}

func CheckConnection() int {
	err := ClientMongo.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
