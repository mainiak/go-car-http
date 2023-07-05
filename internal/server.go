package internal

import (
	"embed"

	"github.com/gin-gonic/gin"
)

//go:embed index.tpl
var embedFS embed.FS

func Serve() {
	data, _ := embedFS.ReadFile("index.tpl")
	print(string(data))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
