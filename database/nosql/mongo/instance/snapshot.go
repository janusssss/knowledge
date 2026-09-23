package main

import (
	"context"
	"fmt"
	"go-career/database/mongo/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client := tools.GetMongoConnect()
	coll := client.Database("instance").Collection("instance_snapshot")
	project_id := "2c47640301b2bbabda6a60fd2c48048c"

	var pipeline mongo.Pipeline
	match := bson.D{{"$match", bson.D{{"project_id", bson.D{{"$eq", project_id}}}}}}
	pipeline = append(pipeline, match)
	group := bson.D{{"$group", bson.D{
		{"_id", "$status"}, // 使用 _id 来指定分组依据的字段
		{"count", bson.D{{"$sum", 1}}},
	}}}
	pipeline = append(pipeline, group)

	cursor, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		panic(err)
	}
	// Prints the average "rating" for each item
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("%v： %v \n", result["_id"], result["count"])
	}
}
