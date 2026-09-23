package main

import (
	"github.com/gin-gonic/gin"
	"go-interview/gin/protoexample/pd"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
			"status":  http.StatusOK,
		})
	})

	router.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Janus"
		msg.Message = "hello world"
		msg.Number = 12
		c.JSON(http.StatusOK, msg)
	})

	router.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "hello world",
			"status":  http.StatusOK,
		})
	})

	router.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "hello world",
			"status":  http.StatusOK,
		})
	})

	router.GET("/someProtobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &pd.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	router.Run(":8080")

}
