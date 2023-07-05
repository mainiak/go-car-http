package internal

import (
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
	fmt.Printf("\n%v\n\n", ipld_storage) // XXX !!!

	root_folder, err := ipld_storage.read_folder(root_cid)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", root_folder) // XXX

	return ipld_storage
}
