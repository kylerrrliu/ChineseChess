# cn-chess-backend

## Setup

Install Go with https://golang.org/doc/install.

Set GOPATH. For example, in ~/.bashrc (or whatever on Mac),
add
```
export GOPATH="~/workdir/"
export GOROOT="/usr/local/go/"
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
workdir can be anything here.

Add in ~/.bashrc for root user:
```
export GOROOT="/usr/local/go"
export PATH=$PATH:$GOROOT/bin
```

Enable the system to find AI lib
```
export LD_LIBRARY_PATH=~/lib
```

#### GoLand setup (if you need an IDE, it's not free though):

Import the repo into GoLand.
To resolve all dependencies:
1. Click Terminal tab at the bottom of the window
2. Run `go env -w GOSUMDB=off`
3. Run `sudo go mod vendor`
4. Run `go mod download` and wait.
5. Restart Goland
6. Goland will prompt about setting up GOROOT and GOPATH. GOROOT should be `/usr/local/go`. Set project GOPATH and global GOPATH to
   `~/workdir`. Check `Index entire GOPATH`.

## Development

Endpoints definition: service.go

GetNextMove function: get_next_move.go

Build: `make`

Run: `make run`

Test: `curl -X POST localhost:7345/cn-chess-backend/get-next-move --data "[[\"aa\", \"bb\"]]"`

