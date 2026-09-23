package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/someJSON", func(c *gin.Context) {
		names := []string{"John", "Doe", "Janus"}
		c.SecureJSON(http.StatusOK, names)
	})
	router.Run(":8080")
}
