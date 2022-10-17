package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	} else {
		fmt.Printf("Ping was successful!\n")
	}

	pinsCollection := client.Database("testing").Collection("pins")

	{
		testPin := bson.D{
			{"href", "https://www.google.com"},
			{"description", "Google"},
			{"extended", ""},
			{"meta", "123456789"},
			{"hash", "123456790"},
			{"time", "string"},
			{"shared", "string"},
			{"toread", "string"},
			{"tags", "string"},
		}

		result, err := pinsCollection.InsertOne(context.TODO(), testPin)
		if err != nil {
			panic(err)
		}

		fmt.Println(result.InsertedID)
	}

	testPins := []interface{}{
		bson.D{{"href", "https://blog.logrocket.com/integrating-mongodb-go-applications/"}, {"description", "Mongo and Go"}},
		bson.D{{"href", "https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-os-x/"}, {"description", "Mongo Manual"}},
	}

	results, err := pinsCollection.InsertMany(context.TODO(), testPins)
	if err != nil {
		panic(err)
	}

	fmt.Println(results.InsertedIDs)

}
