package internal

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-cid"
	carv2 "github.com/ipld/go-car/v2"
	"github.com/ipld/go-car/v2/blockstore"
	"github.com/multiformats/go-multicodec"
)

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

func LoadCAR2(path string, root_cid cid.Cid) *blockstore.ReadOnly {
	robs, err := blockstore.OpenReadOnly(path,
		blockstore.UseWholeCIDs(true),
		carv2.ZeroLengthSectionAsEOF(true),
	)
	if err != nil {
		panic(err)
	}
	defer robs.Close()

	/*
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
	*/

	size, err := robs.GetSize(context.TODO(), root_cid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%v -> %v bytes\n", root_cid, size)

	/*
		block, err := robs.Get(ctx, root_cid)
		if err != nil {
			panic(err)
		}
	*/

	return robs
}

func serve_car(c *gin.Context) {
	c.JSON(200, gin.H{"status": "test"})
}
