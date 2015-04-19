package packet

import (
	"fmt"
)

//PeerPacket implements the Packet interface.
type PeerPacket struct {
	payload       []Bufferable //payload is a slice of bufferable things
	bufferization []byte
}

//CreatePacket takes Bufferable things, and creates a Packet from them.
func (pack *PeerPacket) CreatePacket(things ...Bufferable) (Packet, error) {
	if things == nil {
		return &PeerPacket{}, fmt.Errorf("Bufferable things must be provided.")
	}

	var newPack *PeerPacket
	newPack = new(PeerPacket)
	newPack.payload = things   //place the values in the structure
	err := newPack.bufferize() //create the buffer in the structure
	if err != nil {
		return &PeerPacket{}, fmt.Errorf("Error bufferizing things: %v", err)
	}

	return newPack, nil

}

//Bufferize extracts elements from the fields in the packet struct to
// a single byte slice. This is the complete payload to be shipped through a UDP connection.
func (pack *PeerPacket) bufferize() error {
	//TODO create buffer first, then add bytes (faster for memory allocation)

	buffer := make([]byte, 0) //make an empty buffer that we'll append to

	for ind, bufferable := range pack.payload {
		//iterate through bufferable things to create a megabuffer.
		//this means that the payload itself is a Bufferable

		buf := bufferable.GetBytes()
		if buf == nil {
			return fmt.Errorf("Index %v has a nil buffer.", ind)
		}
		buffer = append(buffer, buf...)
	}

	pack.bufferization = buffer
	return nil
}

func (pack *PeerPacket) GetBuffer() []byte {
	return pack.bufferization
}
