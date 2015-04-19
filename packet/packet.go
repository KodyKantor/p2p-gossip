package packet

import (
	"fmt"
)

//Packet interface defines the functionality for UDP packets.
type Packet interface {
	bufferize() error //converts Bufferable things into a buffer
	CreatePacket(...Bufferable) (Packet, error)
	GetBuffer() []byte
}

//Bufferable interface means that a structure can be placed in a buffer.
type Bufferable interface {
	GetBytes() []byte
}

func init() {
	fmt.Println("Initialed packet package")
}
