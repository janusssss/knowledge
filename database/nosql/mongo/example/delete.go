package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb://localhost:27017/?retryWrites=true&writeConnect=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_guides").Collection("comets")
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"orbitalPeriod", bson.D{{"$gt", 5}}}},
			bson.D{{"orbitalPeriod", bson.D{{"$lt", 85}}}},
		},
		},
	}

	result, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of documents deleted: %d", result.DeletedCount)
}
