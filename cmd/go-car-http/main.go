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
	car_path := os.Args[2]

	/*
	 * `*os.File` supports io.ReadableAt interface needed for CAR factory methods
	 */
	car_fd, err := os.Open(car_path)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := car_fd.Close(); err != nil {
			panic(err)
		}
	}()

	root_cid := internal.ParseCID(cid_str)
	ipld_storage := internal.LoadCAR(car_fd, root_cid)

	internal.Serve(ipld_storage, root_cid)
}
