package internal

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-cid"
)

//go:embed index.tmpl
var embedFS embed.FS

// Gin middleware
func root_cid_mw(ipld_storage *IPLD_Storage, root_cid cid.Cid) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Pass storage
		c.Set("ipld_storage", ipld_storage)
		c.Set("root_cid", root_cid)

		// Process the request
		c.Next()
	}
}

func Serve(ipld_storage *IPLD_Storage, root_cid cid.Cid) {
	index_tmpl, _ := embedFS.ReadFile("index.tmpl")

	mr := multitemplate.NewRenderer()
	mr.AddFromString("index", string(index_tmpl))

	r := gin.Default()

	// https://gin-gonic.com/docs/examples/html-rendering/#custom-template-renderer
	r.HTMLRender = mr

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"Title": "go-car-http",
			"H1":    "Demo go-car-http",
			//"BodyP": fmt.Sprintf("CID: %s, CAR file: %s\n", root_cid, car_str),
			"BodyP": fmt.Sprintf("CID: %s\n", root_cid),
		})
	})

	// TODO: "/info"

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	rootURL := r.Group("/root")
	rootURL.Use(root_cid_mw(ipld_storage, root_cid))
	{
		rootURL.GET("/", serve_car)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
