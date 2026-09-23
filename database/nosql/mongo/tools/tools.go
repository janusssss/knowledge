package tools

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var once sync.Once
var client *mongo.Client

func GetMongoConnect() *mongo.Client {
	once.Do(func() {
		uri := "mongodb://localhost:27017/?retryWrites=true&writeConcern=majority"
		var err error
		client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
	})
	return client
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
