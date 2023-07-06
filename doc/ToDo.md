# To Do

## Known issues

- Folder currently can't have `/` at end of URI. It should be optional and work with or without it at end of URI.
- Only RAW files works, DAG-PB files does not work.

## To be done

- [ ] DAG-PB can be folder or *file* (!!)
- [ ] Support dag-json to just return JSON (~~update demo.car~~)
- [ ] Add logging library to replace `fmt.Printf` calls
- [.] Better error handling

## Nice to have

- [ ] MIME support
- [ ] Use streams when serving files

## Done

- [x] Test sub-folders (~~update demo.car~~)
- [x] Fix issue with `storage.Has()` CID test call
