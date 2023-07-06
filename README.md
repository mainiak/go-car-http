# go-car-http

## About

_WORK IN PROGRESS_

Plain HTTP server for CAR files.  
Crude implementation as Proof of Concept - not for production.  
Use at your own risk.

## Dev

```
go build -o build ./... && \
./build/go-car-http \
bafybeih3eahbwomynpl3pgpzl6dwkgvyishjwlx77wupujo4kacikciosq \
data/demo_folder.car
```

## Install

```
go install github.com/mainiak/go-car-http/cmd/go-car-http@latest
```

## Use

```
go-car-http <CID> <file.car>
open http://localhost:8080/root
```

## Uninstall

```
rm -v $(which go-car-http)
```
