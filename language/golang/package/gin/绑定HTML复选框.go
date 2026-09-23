package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("gin/views/*")
	r.GET("/", indexHandler)
	r.POST("/", formHandler)
	r.Run()
}

type myForm struct {
	Colors []string `form:"colors[]"`
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{"colors": fakeForm.Colors})
}
