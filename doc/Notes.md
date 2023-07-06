# Notes

## Features

- Serve content from within the CAR file (that's it).
- __No__ intention whatsoever to add __SSL__. Please use Traefik/Ngnix/etc. for SSL in front of `go-car-http`.

## Use cases

### Demo

`go-car-http` can serve any static conten from within the CAR file.  
This is a Proof of Concept use case.

### Linux package mirrors

To serve content from IPFS - by using CAR files - and to create Linux HTTP mirrors (see [1], [2]).  
To share CAR file over IPFS public network/swarm, please use Frisbii (see [6]) for same CAR file.

This is temporary solution until projects like WNFS (see [3]),  
or local mount with IPFS Kubo/Desktop is stable/mature/fast enough to be used instead.

You can also checkout existing IPFS Cluster based solution (see [4]).

## CAR files

### Creating CAR file

```
ipfs add -r folder/
ipfs dag export <cid> >file.car

## using `car` command (??)
```

### CAR CLI

```
go install github.com/ipld/go-car/cmd/car@latest

car inspect data/demo_folder.car
car verify data/demo_folder.car || echo "failed"
```

See [5] for more.

### Frisbii

To expose data to public IPFS network/swarm - like ie. your CAR file with Linux packages

```
go install github.com/ipld/frisbii/cmd/frisbii@latest
frisbii --announce=roots --car=data/demo_folder.car
```

See [6] for more.

### Lassie

Download content from IPFS network/swarm so you can serve it in your local network/computer with `go-car-http`.

```
go install github.com/filecoin-project/lassie/cmd/lassie@latest
lassie -o file.car <CID>[/path/to/content]
```

## Links

- [1] https://wiki.alpinelinux.org/wiki/How_to_setup_a_Alpine_Linux_mirror
- [2] https://wiki.archlinux.org/title/mirrors
- [3] https://guide.fission.codes/developers/webnative/file-system-wnfs
- [4] https://github.com/RubenKelevra/pacman.store
- [5] https://github.com/ipld/go-car
- [6] https://github.com/ipld/frisbii
- [7] https://github.com/filecoin-project/lassie
