package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-cid"
)

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
	r := gin.Default()
	r.HTMLRender = get_templates()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"Title": "go-car-http",
			"H1":    "Demo go-car-http",
			//"BodyP": fmt.Sprintf("CID: %s, CAR file: %s\n", root_cid, car_str),
			"BodyP": fmt.Sprintf("CID: %s\n", root_cid),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	rootURL := r.Group("/root")
	rootURL.Use(root_cid_mw(ipld_storage, root_cid))
	{
		rootURL.GET("/", serve_root)
		// TODO: support any URL bellow/under '/root' -> path parsing
	}

	filesURL := r.Group("/files")
	filesURL.Use(root_cid_mw(ipld_storage, root_cid))
	{
		filesURL.GET("/", serve_files)
	}

	infoURL := r.Group("/info")
	infoURL.Use(root_cid_mw(ipld_storage, root_cid))
	{
		infoURL.GET("/", serve_info)
		// TODO: support any URL bellow/under '/info' -> path parsing
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
