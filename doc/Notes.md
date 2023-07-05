# Notes

## To Do (someday)

- Fix issue with `storage.Has()` CID test call
- Better error handling
- Add logging library to replace `fmt.Printf` calls

## CAR files

### CAR CLI

```
go install github.com/ipld/go-car/cmd/car@latest

car inspect data/demo_folder.car
car verify data/demo_folder.car || echo "failed"
```

### Frisbii

To expose data to public IPFS network/swarm

```
go install github.com/ipld/frisbii/cmd/frisbii@latest
frisbii --announce=roots --car=data/demo_folder.car
```

## Links

- https://github.com/ipld/frisbii
