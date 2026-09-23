package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type requestRecord struct {
	LastRequestTime time.Time
}

var (
	requestCache = make(map[string]*requestRecord)
	mutex        sync.Mutex
)

func main() {
	router := gin.Default()
	router.GET("", Handler)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func Handler(c *gin.Context) {
	clientId := c.ClientIP()
	mutex.Lock()
	record, exists := requestCache[clientId]
	if !exists {
		record = &requestRecord{}
		requestCache[clientId] = record
	}
	mutex.Unlock()

	currentTime := time.Now()
	if record.LastRequestTime.IsZero() || currentTime.Sub(record.LastRequestTime) > 2*time.Minute {
		mutex.Lock()
		record.LastRequestTime = currentTime
		mutex.Unlock()
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	} else {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "重复请求，请稍后再试",
		})
	}
}
