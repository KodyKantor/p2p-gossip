//Package packet provides tools to turn data into a buffer which can be sent through a packet.
package packet

import (
	"fmt"
)

//Packet interface defines the functionality for UDP packets.
type Packet interface {
	bufferize() error //converts Bufferable things into a buffer
	CreatePacket(...Bufferable) (Packet, error)
	GetBufferization() []byte
	CreatePacketFromBytes([]byte) (Packet, error)
}

//Bufferable interface means that a structure can be placed in a buffer.
type Bufferable interface {
	GetBytes() []byte
}

func init() {
	fmt.Println("Initialed packet package")
}
