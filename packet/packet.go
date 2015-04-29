//Package packet provides tools to turn data into a buffer which can be sent through a packet.
package packet

import "github.com/Sirupsen/logrus"

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
	logrus.Debugln("Initialed packet package")
}

//BufferizableString is a wrapper around the string primitive
//so that it can be easily placed in a buffer to be sent in a packet.
type BufferizableString struct {
	Str string
}

//GetBytes will convert a string into its byte-slice representation.
func (s *BufferizableString) GetBytes() []byte {
	return []byte(s.Str)
}
