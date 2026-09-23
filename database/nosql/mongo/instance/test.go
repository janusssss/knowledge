package main

import (
	"context"
	"fmt"
	"go-career/database/mongo/tools"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	client := tools.GetMongoConnect()
	coll := client.Database("instance").Collection("instance_snapshot")
	result := coll.FindOne(context.TODO(), bson.D{{"status", "AVAILABLE"}})
	raw, _ := result.Raw()
	elements, _ := raw.Elements()
	for _, v := range elements {
		fmt.Printf("%s: %s\n", v.Key(), v.Value())
	}
}
