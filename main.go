package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// connectionString := "mongodb://localhost:27017/?replicaSet=rs0"
	connectionString := "mongodb://mongo-primary:27017,mongo-secondary:27017"
	cOpts := options.Client().
		ApplyURI(connectionString).
		SetReplicaSet("testReplica").
		SetConnectTimeout(2 * time.Second).
		SetServerSelectionTimeout(2 * time.Second)

	client, err := mongo.NewClient(cOpts)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")

	defer client.Disconnect(ctx)

	db := client.Database("test")
	usersCollection := db.Collection("users")
	usersStream, err := usersCollection.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		log.Println(err)
	}
	if usersStream == nil {
		log.Fatalln("unable to watch stream")
	}

	log.Println("watching for updates")

	for usersStream.Next(context.TODO()) {
		var data bson.M
		if err := usersStream.Decode(&data); err != nil {
			panic(err)
		}
		log.Printf("got an update: %v\n", data)
	}
}
