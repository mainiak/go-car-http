package internal

import (
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car/v2/storage"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/linking"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
)

type IPLD_Storage struct {
	CARR storage.ReadableCar
	LCtx linking.LinkContext
	LSys linking.LinkSystem
}

func NewIPLD_Storage(carr storage.ReadableCar, lctx linking.LinkContext, lsys linking.LinkSystem) *IPLD_Storage {
	return &IPLD_Storage{
		CARR: carr,
		LCtx: lctx,
		LSys: lsys,
	}
}

func (is *IPLD_Storage) read_folder(folder_cid cid.Cid) (map[string]datamodel.Link, error) {
	var folder_lnk datamodel.Link
	folder := make(map[string]datamodel.Link)

	folder_lnk = cidlink.Link{
		folder_cid,
	}

	//np := basicnode.Prototype.Any // NOPE; it works, but not what we need
	np := dagpb.Type.PBNode
	lnk := datamodel.Link(folder_lnk)

	node, err := is.LSys.Load(is.LCtx, lnk, np)
	if err != nil {
		return nil, err
	}

	pbnode := node.(dagpb.PBNode)
	//fmt.Printf("%v\n", pbnode.Links) // XXX
	//fmt.Printf("%v\n", pbnode.Data) // XXX

	links := pbnode.Links.Iterator()
	for i := 0; i < int(pbnode.Links.Length()); i++ {
		_, pb_link := links.Next()
		file_name, file_cid, _ := get_pblink(pb_link)
		folder[file_name] = file_cid
	}

	return folder, nil
}
