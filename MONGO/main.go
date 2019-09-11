package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ListAndReview ... strct for collection in DB
type ListAndReview struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	listing_url string
	name        string `json:"name,omitempty" bson:"name,omitempty"`
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb+srv://user:pass@ashucluster-hwkoc.mongodb.net/test?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	var result ListAndReview
	filter := bson.D{}
	collection := client.Database("sample_airbnb").Collection("listingsAndReviews")

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}
