package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-cid"
)

// Gin middleware
func ipld_mw(ipld_storage *IPLD_Storage, root_cid cid.Cid) gin.HandlerFunc {
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

	// Enabling for all routes
	r.Use(ipld_mw(ipld_storage, root_cid))

	r.GET("/", serve_root)

	apiURL := r.Group("/_")
	{
		apiURL.GET("/about", serve_about)
		apiURL.GET("/files", serve_files)
		apiURL.GET("/info", serve_info)
	}

	r.NoRoute(func(c *gin.Context) {
		/*
			// XXX
			//full_path := c.FullPath()
			//fmt.Printf("# full_path: '%s'\n", full_path)
			// XXX
			request_uri := c.Request.RequestURI
			//fmt.Printf("# uri: '%s'\n", request_uri) // XXX

			if strings.HasPrefix(request_uri, "/root/") {
				serve_subroot(c)
			} else {
				serve_not_found(c)
			}
		*/

		serve_subroot(c)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
