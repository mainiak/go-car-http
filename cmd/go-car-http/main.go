package main

import (
	"fmt"
	"os"

	"github.com/mainiak/go-car-http/internal"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("%s <cid> <file.car>", os.Args[0])
	}

	cid_str := os.Args[1]
	car_str := os.Args[2]

	root_cid := internal.ParseCID(cid_str)

	internal.LoadCAR(car_str, root_cid)

	//internal.Serve(bs, root_cid)
}
