[![wercker status](https://app.wercker.com/status/a6d202a73b8cae2aa7dd3b1da1ef3c12/m "wercker status")](https://app.wercker.com/project/bykey/a6d202a73b8cae2aa7dd3b1da1ef3c12)
[![GoDoc](https://godoc.org/github.com/KodyKantor/p2p-gossip?status.svg)](https://godoc.org/github.com/KodyKantor/p2p-gossip)

# p2p-gossip
A p2p gossip protocol for requesting artifacts.

## Building
From the top-level directory:
<br />
`go get -v` <br />
`go build main.go`

## Running
Run the peer by passing in the `-p` argument to the binary. This will run the peer in the foreground.

For example, `./gossip -p --verbose`

Instantiate the binary subsequent times to receive resources with the `get` argument.

For example, `./gossip get cats.jpg`

## Running Tests
Running the unit tests doesn't require any setup.
`go test ./...`

## Package Descriptions
### ID
The ID package provides functionality for generating unique, random IDs. IDs consist of a simple byte slice, and a size.
<br />
To use the ID package, I recommend making a 'master' ID running new(PeerID), and then calling the ServeIDs function on that ID pointer.
The ServeIDs function will feed new IDs through a channel (thread safe queue primitive for inter-process communication).
<br /><br />

### TTL
The TTL package implements a time-to-live structure. This is a simple wrapper around an integer that represents a
time-to-live for a packet.

### UDP
The UDP package provides the functionality to communicate between peers in the peer network.

### Packet
The Packet package transforms the types that need to be sent in a packet to a single buffer.
