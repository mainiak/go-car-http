# Notes

## To Do (someday)

- [ ] Stream (?) raw codec objects as binary
- [ ] Support dag-json to just return JSON (update demo.car)
- [ ] Test sub-folders (update demo.car)
- [ ] Fix issue with `storage.Has()` CID test call
- [.] Better error handling
- [ ] Add logging library to replace `fmt.Printf` calls

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

### Frisbii

To expose data to public IPFS network/swarm

```
go install github.com/ipld/frisbii/cmd/frisbii@latest
frisbii --announce=roots --car=data/demo_folder.car
```

## Links

- https://github.com/ipld/frisbii
