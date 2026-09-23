package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html = template.Must(template.New("http").Parse(`
<html>
<head>
<title>Https Test</title>
<script src="/assets/app.js"></script>
</head>
<body>
<h1 style="color:red;">Welcome, Ginner</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)
	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Println("push err:", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "ok",
		})
	})
	r.RunTLS(":8080", "./testdate/server.pem", "./testdate/server.key")
}
