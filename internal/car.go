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

func print_pblink(pblink dagpb.PBLink) {
	//fmt.Printf("%v // %s\n", pblink, pblink.Kind()) // XXX

	if pblink.Name.Exists() {
		//fmt.Printf("Name: %s\n", pblink.Name) // XXX
		name_str := pblink.Name.Must()
		fmt.Printf("# name: %s\n", name_str)
	}

	//fmt.Printf("Hash: %s\n", pblink.Hash) // XXX
	if !pblink.Hash.IsAbsent() && !pblink.Hash.IsNull() {
		lnk := pblink.Hash.Link()
		fmt.Printf("# link: %s\n", lnk)
	}

	//fmt.Printf("Tsize: %s\n", pblink.Tsize) // XXX
	if pblink.Tsize.Exists() {
		size := pblink.Tsize.Must().Int()
		fmt.Printf("# size: %d\n", size)
	}

	/*
		if pblink.Kind() == datamodel.Kind_Map {
			mi := pblink.MapIterator()
			mi_len := int(pblink.Length())
			for j := 0; j < mi_len; j++ {
				k, v, e := mi.Next()
				fmt.Printf("%s %s %s", k, v, e)
			}
		}
	*/
}

func LoadCAR(path string, root_cid cid.Cid) {
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

	rcar.Index() // TODO: needed?
	roots := rcar.Roots()

	//fmt.Printf("%v\n", roots) // XXX
	fmt.Printf("\nRoots: \n")
	for _, rcid := range roots {
		fmt.Printf(" - %v\n", rcid)
	}
	fmt.Printf("\n")

	lctx := linking.LinkContext{}
	lsys := cidlink.DefaultLinkSystem()
	lsys.SetReadStorage(rcar)

	//np := basicnode.Prototype.Any // NOPE; it works, but not what we need
	np := dagpb.Type.PBNode
	lnk := datamodel.Link(root_lnk)

	node, err := lsys.Load(lctx, lnk, np)
	if err != nil {
		panic(err)
	}

	/*
		// XXX
		fmt.Printf("%v\n", node)
		fmt.Printf("we loaded a %s with %d entries\n", node.Kind(), node.Length())
		// XXX
	*/

	pbnode := node.(dagpb.PBNode)
	/*
		// XXX
		fmt.Printf("%v\n", pbnode)
		fmt.Printf("we loaded a %s / %s with %d entries\n", pbnode.Kind(), pbnode.Type(), pbnode.Length())
		// XXX
	*/

	//fmt.Printf("%v\n", pbnode.Links) // XXX
	//fmt.Printf("%v\n", pbnode.Data) // XXX

	links := pbnode.Links.Iterator()
	for i := 0; i < int(pbnode.Links.Length()); i++ {
		//idx, pb_link := links.Next()
		//fmt.Printf("idx: %d, ", idx) // no newline on purpose - XXX
		_, pb_link := links.Next()
		print_pblink(pb_link)
	}

	//return lsys
}

func serve_car(c *gin.Context) {
	c.JSON(200, gin.H{"status": "test"})
}
