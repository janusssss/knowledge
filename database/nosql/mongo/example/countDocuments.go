package main

import (
	"context"
	"go-career/database/mongo/tools"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	c := tools.GetMongoConnect()
	count, err := c.Database("instance").Collection("instance_snapshot").CountDocuments(context.TODO(), bson.D{})
	tools.CheckErr(err)
	println(count)
}
