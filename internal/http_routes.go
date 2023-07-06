package internal

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multicodec"
)

func req_root_cid(c *gin.Context) cid.Cid {
	obj, exists := c.Get("root_cid")

	if !exists {
		serve_error("root_cid not found", nil, c)
		//c.Abort()
		return cid.Undef // TODO: Correct error/nil type?
	}

	root_cid := obj.(cid.Cid)
	return root_cid
}

func req_ipld_storage(c *gin.Context) *IPLD_Storage {
	obj, exists := c.Get("ipld_storage")

	if !exists {
		serve_error("ipld_storage not found", nil, c)
		//c.Abort()
		return nil
	}

	ipld_storage := obj.(*IPLD_Storage)
	return ipld_storage
}

func serve_index(c *gin.Context) {
	root_cid := req_root_cid(c)
	if root_cid == cid.Undef {
		return
	}

	c.HTML(http.StatusOK, "index", gin.H{
		"Title": "go-car-http",
		"H1":    "Demo go-car-http",
		//"BodyP": fmt.Sprintf("CID: %s, CAR file: %s\n", root_cid, car_str), // XXX
		"BodyP": fmt.Sprintf("CID: %s\n", root_cid),
	})
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
		serve_error("failed reading root folder", err, c)
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
		serve_error("failed reading root folder", err, c)
		return
	}

	fmt.Printf("%v\n", folder) // XXX
	//c.JSON(200, gin.H{"files": folder}) // XXX

	c.HTML(http.StatusOK, "files", gin.H{
		"files":      folder,
		"full_path":  c.FullPath(),
		"folder_cid": root_cid.String(),
	})
}

func serve_error(msg string, err error, c *gin.Context) {
	fmt.Printf("ERR: %s, reason: %s\n", msg, err)
	c.JSON(500, gin.H{
		"error":  msg,
		"reason": err,
	})
	//c.AbortWithError(500, err)
	c.Abort()
}

func serve_not_found(c *gin.Context) {
	request_uri := c.Request.RequestURI
	c.JSON(404, gin.H{
		"code":        "PAGE_NOT_FOUND",
		"message":     "Page not found",
		"request_uri": request_uri,
	})
}

// URL '/root/*'
func serve_subroot(c *gin.Context) {
	request_uri := c.Request.RequestURI
	uri_path_orig := strings.Split(request_uri, "/")[2:]
	fmt.Printf("uri_path: %v\n", uri_path_orig) // XXX
	//fmt.Printf("full_path: %s\n", c.FullPath()) // XXX - not working when handling "not found" requests

	root_cid := req_root_cid(c)
	if root_cid == cid.Undef {
		return
	}

	ipld_storage := req_ipld_storage(c)

	file_name := uri_path_orig[0]
	uri_path := uri_path_orig[1:]
	folder_cid := root_cid

	// TODO: recurision point starts here
	// TODO: recursion until len(uri_path) == 0 ... or RAW IPLD object found

	do_repeat := true
	for do_repeat {

		do_repeat, folder_cid = walk_folder(c, ipld_storage, folder_cid, file_name)

		if do_repeat && len(uri_path) > 0 {
			// walk uri_path
			file_name, uri_path = uri_path[0], uri_path[1:]
			fmt.Printf("## file_name: %s\n## uri_path: %v\n", file_name, uri_path) // XXX

			continue
		}

		if do_repeat && len(uri_path) == 0 {
			folder, err := ipld_storage.read_folder(folder_cid)
			if err != nil {
				serve_error("failed reading root folder", err, c)
				return
			}

			fmt.Printf("%v\n", folder) // XXX
			//c.JSON(200, gin.H{"files": folder}) // XXX

			c.HTML(http.StatusOK, "files", gin.H{
				"files":      folder,
				"full_path":  fmt.Sprintf("%s/", request_uri),
				"folder_cid": folder_cid.String(),
			})
			break
		}
	}
}

func walk_folder(c *gin.Context, ipld_storage *IPLD_Storage, folder_cid cid.Cid, file_name string) (bool, cid.Cid) {
	folder, err := ipld_storage.read_folder(folder_cid)
	if err != nil {
		serve_error("failed reading root folder", err, c)
		return false, cid.Undef
	}

	obj_lnk, obj_present := folder[file_name]
	if !obj_present {
		serve_error("File not found", fmt.Errorf("No such file: %s", file_name), c)
		return false, cid.Undef
	}

	lnk_str := obj_lnk.String()
	fmt.Printf("lnk: '%s'\n", lnk_str) // XXX
	obj_cid := ParseCID(lnk_str)
	obj_codec := multicodec.Code(obj_cid.Prefix().Codec)
	fmt.Printf("Codec: 0x%04x (%s)\n", uint64(obj_codec), obj_codec)

	/*
		// TODO: Convert to structure and serve as JSON
		if obj_codec == multicodec.DagJson {
		}
	*/

	// Process 'raw' IPLD objects
	if obj_codec == multicodec.Raw {
		//ipld_storage.CARR.GetStream(context.TODO(), lnk_str)
		//c.DataFromReader()

		data, err := ipld_storage.CARR.Get(context.TODO(), obj_lnk.Binary())
		if err != nil {
			serve_error("Failed to read from CAR file", err, c)
			return false, cid.Undef
		}

		c.Data(http.StatusOK, "application/octet-stream", data)
		return false, cid.Undef
	}

	if obj_codec == multicodec.DagPb {
		return true, obj_cid // yes to recursion
	}

	err = fmt.Errorf("codec: 0x%04x (%s)", uint64(obj_codec), obj_codec)
	serve_error("Codec not supported", err, c)
	//c.Abor()

	return false, cid.Undef
}
