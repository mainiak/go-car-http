package internal

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multicodec"
)

func ParseCID(cid_str string) cid.Cid {
	fmt.Printf("\ncid: %s\n", cid_str)
	obj_cid, err := cid.Decode(cid_str)
	if err != nil {
		panic(err)
	}

	obj_prefix := obj_cid.Prefix()
	fmt.Printf("Version: %d\n", obj_prefix.Version)
	fmt.Printf("Codec: 0x%04x (%s)\n", obj_prefix.Codec, multicodec.Code(obj_prefix.Codec))
	fmt.Printf("MhType: 0x%02x (%s)\n", obj_prefix.MhType, multicodec.Code(obj_prefix.MhType))
	fmt.Printf("MhLength: %d\n", obj_prefix.MhLength)

	return obj_cid
}
