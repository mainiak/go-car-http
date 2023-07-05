package internal

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car/v2/storage"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/linking"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
)

/*
func LoadCAR(path string, asked_root_cid cid.Cid) {
	fmt.Println("\nFile:", path)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	br, err := carv2.NewBlockReader(f)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("CAR Version: %d\n", br.Version)
	fmt.Printf("Roots:\n")

	match_root_cid := false
	for _, root_cid := range br.Roots {
		fmt.Printf(" - %s\n", root_cid)

		if asked_root_cid == root_cid {
			match_root_cid = true

			cid_codec := root_cid.Prefix().Codec
			code := multicodec.Code(cid_codec)
			fmt.Printf("Codec: 0x%04x (%s)\n", cid_codec, code)

			if code != multicodec.DagPb {
				// panic will f.Close() ... because defer
				panic(fmt.Errorf("CID: %s (%s) is not a DAG-PB", root_cid, code))
			}
		}
	}

	if !match_root_cid {
		panic(fmt.Errorf("Root CID (%s) not found", asked_root_cid))
	}

	fmt.Println("First 5 block CIDs:")

	for i := 0; i < 5; i++ {
		bl, err := br.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		c := bl.Cid()
		fmt.Printf(" - %v\n", c)
	}

	for _, root_cid := range br.Roots {
		ParseCID(root_cid.String())
	}
}
*/

func LoadCAR3(path string, root_cid cid.Cid) {
	fmt.Printf("%v\n\n", dagpb.Type) // FIXME

	var root_lnk datamodel.Link
	root_lnk = cidlink.Link{
		root_cid,
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	rcar, err := storage.OpenReadable(f)
	if err != nil {
		panic(err)
	}

	rcar.Index()
	roots := rcar.Roots()

	fmt.Printf("%v\n", roots) // XXX
	fmt.Printf("Roots3: \n")
	for _, rcid := range roots {
		fmt.Printf(" - %v\n", rcid)
	}

	lctx := linking.LinkContext{}
	lsys := cidlink.DefaultLinkSystem()
	lsys.SetReadStorage(rcar)

	//np := basicnode.Prototype.Any
	np := dagpb.Type.PBNode
	lnk := datamodel.Link(root_lnk)

	node, err := lsys.Load(lctx, lnk, np)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", node)
	fmt.Printf("we loaded a %s with %d entries\n", node.Kind(), node.Length())

	pbnode := node.(dagpb.PBNode)
	fmt.Printf("%v\n", pbnode)
	fmt.Printf("we loaded a %s / %s with %d entries\n", pbnode.Kind(), pbnode.Type(), pbnode.Length())

	fmt.Printf("%v\n", pbnode.Links)
	fmt.Printf("%v\n", pbnode.Data)

	links := pbnode.Links.Iterator()
	for i := 0; i < int(pbnode.Links.Length()); i++ {
		idx, pb_link := links.Next()
		fmt.Printf("%d: %v // %s\n", idx, pb_link, pb_link.Kind())

		if pb_link.Kind() == datamodel.Kind_Map {
			mi := pb_link.MapIterator()
			mi_len := int(pb_link.Length())
			for j := 0; j < mi_len; j++ {
				k, v, e := mi.Next()
				fmt.Printf("%s %s %s", k, v, e)
			}
		}
	}

	/*
		mi := node.MapIterator()
		for i := 0; i < int(node.Length()); i++ {
			key, val, err := mi.Next()
			if err != nil {
				panic(err)
			}
			if key == nil {
				break
			}
			fmt.Printf("%s, %d, %s\n", key.Kind(), key.Length(), key)
			fmt.Printf("%s, %d, %s\n", val.Kind(), val.Length(), val)

			str1, err := key.AsString()
			fmt.Printf("%s // %s\n", str1, err)

			str2, err := val.AsString()
			fmt.Printf("%s // %s\n", str2, err)
		}
	*/

	//return lsys
}

func serve_car(c *gin.Context) {
	c.JSON(200, gin.H{"status": "test"})
}
