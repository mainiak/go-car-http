package internal

import (
	"embed"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

//go:embed index.tmpl
var embedFS embed.FS

func Serve() {
	index_tmpl, _ := embedFS.ReadFile("index.tmpl")

	mr := multitemplate.NewRenderer()
	mr.AddFromString("index", string(index_tmpl))

	r := gin.Default()

	// https://gin-gonic.com/docs/examples/html-rendering/#custom-template-renderer
	r.HTMLRender = mr

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"Title": "go-car-http",
			"H1":    "go-car-http",
			"BodyP": "Demo go-car-http.",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
