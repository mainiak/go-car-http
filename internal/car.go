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

func get_pblink(pblink dagpb.PBLink) (string, datamodel.Link, int64) {
	var name_str string
	var lnk datamodel.Link
	var size int64

	//fmt.Printf("%v // %s\n", pblink, pblink.Kind()) // XXX

	if pblink.Name.Exists() {
		//fmt.Printf("Name: %s\n", pblink.Name) // XXX
		name_str = pblink.Name.Must().String()
		fmt.Printf("# name: %s\n", name_str)
	}

	//fmt.Printf("Hash: %s\n", pblink.Hash) // XXX
	if !pblink.Hash.IsAbsent() && !pblink.Hash.IsNull() {
		lnk = pblink.Hash.Link()
		fmt.Printf("# link: %s\n", lnk)
	}

	//fmt.Printf("Tsize: %s\n", pblink.Tsize) // XXX
	if pblink.Tsize.Exists() {
		size = pblink.Tsize.Must().Int()
		fmt.Printf("# size: %d\n", size)
	}

	fmt.Printf("[%s, %s, %d]\n", name_str, lnk, size)

	return name_str, lnk, size
}

func LoadCAR(path string, root_cid cid.Cid) *IPLD_Storage {
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

	/*
		// FIXME
		fmt.Printf("CID: %s\n", root_cid) // XXX
		ParseCID(root_cid.String()) // XXX
		cid_exists, err := rcar.Has(context.TODO(), root_cid.String())
		if err != nil {
			// WTH ?!
			// CID: bafybeifg2xqiapzayizkm652rntpqlq5dxxnbq3u7k7uim3dyvwq7cuugy
			//panic: bad CID key: invalid cid: invalid cid: expected 1 as the cid version number, got: 98
			panic(err)
		}
		if !cid_exists {
			panic(fmt.Errorf("Your requested CID: %s, is not in CAR file.", root_cid))
		}
		// FIXME
	*/

	//fmt.Printf("%v\n", roots) // XXX
	fmt.Printf("\nRoots: \n")
	roots := rcar.Roots()
	for _, rcid := range roots {
		fmt.Printf(" - %v\n", rcid)
	}
	fmt.Printf("\n")

	lctx := linking.LinkContext{}
	lsys := cidlink.DefaultLinkSystem()
	lsys.SetReadStorage(rcar)

	ipld_storage := NewIPLD_Storage(rcar, lctx, lsys)

	root_folder, err := ipld_storage.read_folder(root_cid)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", root_folder) // XXX

	return ipld_storage
}

func serve_car(c *gin.Context) {
	c.JSON(200, gin.H{"status": "test"})
}
