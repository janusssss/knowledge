package main

import (
	"context"
	"fmt"
	"go-career/database/mongo/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	client := tools.GetMongoConnect()
	coll := client.Database("instance").Collection("instance_snapshot")

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"status", "AVAILABLE"},
			{"project_id", "2c47640301b2bbabda6a60fd2c48048c"},
		},
		}},

		{{"$project", bson.D{
			{"yearMonth", bson.D{ // 创建一个新的字段 yearMonth 用于分组
				{"$dateToString", bson.D{
					{"format", "%Y-%m"}, // 格式化为 "YYYY-MM"
					{"date", bson.D{
						{"$toDate", "$updated_at"}, // 将 created_at 转换为日期
					}},
				}},
			}},
		}}},

		{{"$group", bson.D{
			{"_id", "$yearMonth"},          // 按照 yearMonth 分组
			{"count", bson.D{{"$sum", 1}}}, // 统计每个分组的数量
		}}},

		{{"$sort", bson.D{
			{"_id", 1}, // 按 _id (即年月) 升序排序
		}}},
	}

	cursor, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatalf("聚合操作失败: %v", err)
	}
	defer cursor.Close(context.TODO())

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatalf("读取结果失败: %v", err)
	}

	// 打印结果
	for _, result := range results {
		fmt.Printf("时间: %s, 数量: %d\n", result["_id"], result["count"])
	}
}
