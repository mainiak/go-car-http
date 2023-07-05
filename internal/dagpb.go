package internal

import (
	"fmt"

	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime/datamodel"
)

func get_pblink(pblink dagpb.PBLink) (string, datamodel.Link, int64) {
	var name_str string
	var lnk datamodel.Link // CID
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
