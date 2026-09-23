package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/testing", startPage)
	r.Run()
}

type person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func startPage(c *gin.Context) {
	var p person
	if c.ShouldBindQuery(&p) == nil {
		log.Println(p.Name)
		log.Println(p.Address)
		log.Println(p.Birthday)
	}
	c.String(http.StatusOK, "Hello, World!")
}
