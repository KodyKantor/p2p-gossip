# p2p-gossip
A p2p gossip protocol for requesting artifacts.

## Building
From the top-level directory:
`go build main.go`

## Running Tests
Navigate to a sub-directory
`go test`

## Package Descriptions
### ID
The ID package provides functionality for generating unique, random IDs. IDs consist of a simple byte slice, and a size.
<br />
To use the ID package, I recommend making a 'master' ID using the New() function, and then calling the ServeIDs function on that ID.
The ServeIDs function will feed new IDs through a channel (thread safe queue primitive for inter-process communication).
<br /><br />

### TTL
The TTL package implements a time-to-live structure. This is a simple wrapper around an integer that represents a
time-to-live for a packet.

### UDP
The UDP package provides the functionality to communicate between peers in the peer network.
