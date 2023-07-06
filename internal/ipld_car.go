package internal

import (
	"context"
	"fmt"
	"io"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car/v2/storage"
	"github.com/ipld/go-ipld-prime/linking"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
)

func LoadCAR(car_fd io.ReaderAt, root_cid cid.Cid) *IPLD_Storage {
	rcar, err := storage.OpenReadable(car_fd)
	if err != nil {
		panic(err)
	}

	rcar.Index() // TODO: needed?

	/*
		fmt.Printf("CID: %s\n", root_cid) // XXX
		ParseCID(root_cid.String())       // XXX
	*/
	root_cid_lnk := cidlink.Link{root_cid}
	cid_exists, err := rcar.Has(context.TODO(), root_cid_lnk.Binary())
	if err != nil {
		panic(err)
	}
	if !cid_exists {
		panic(fmt.Errorf("Your requested CID: %s, is not in CAR file.", root_cid))
	}

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
