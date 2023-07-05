package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-cid"
)

func req_root_cid(c *gin.Context) cid.Cid {
	obj, exists := c.Get("root_cid")

	if !exists {
		c.JSON(500, gin.H{
			"error": "root_cid not found",
		})
		return cid.Undef // TODO: Correct error/nil type?
	}

	root_cid := obj.(cid.Cid)
	return root_cid
}

func req_ipld_storage(c *gin.Context) *IPLD_Storage {
	obj, exists := c.Get("ipld_storage")

	if !exists {
		c.JSON(500, gin.H{
			"error": "ipld_storage not found",
		})
		return nil
	}

	ipld_storage := obj.(*IPLD_Storage)
	return ipld_storage
}

// URL '/info/'
func serve_info(c *gin.Context) {
	root_cid := req_root_cid(c)
	if root_cid == cid.Undef {
		return
	}

	c.JSON(200, gin.H{
		"root_cid": root_cid.String(),
	})
}

// URL '/files/'
func serve_files(c *gin.Context) {
	root_cid := req_root_cid(c)
	if root_cid == cid.Undef {
		return
	}

	ipld_storage := req_ipld_storage(c)
	folder, err := ipld_storage.read_folder(root_cid)
	if err != nil {
		c.JSON(500, gin.H{
			"error":  "failed reading root folder",
			"reason": err,
		})
		return
	}

	c.JSON(200, gin.H{"files": folder})
}

// URL '/root/'
func serve_root(c *gin.Context) {
	root_cid := req_root_cid(c)
	if root_cid == cid.Undef {
		return
	}

	ipld_storage := req_ipld_storage(c)
	folder, err := ipld_storage.read_folder(root_cid)
	if err != nil {
		c.JSON(500, gin.H{
			"error":  "failed reading root folder",
			"reason": err,
		})
		return
	}

	// FIXME
	c.JSON(200, gin.H{"files": folder})
}
