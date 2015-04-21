package packet

import (
	"fmt"

	"github.com/kodykantor/p2p-gossip/id"
	"github.com/kodykantor/p2p-gossip/ttl"
)

//PeerPacket implements the Packet interface.
type PeerPacket struct {
	payload       []Bufferable //payload is a slice of bufferable things
	bufferization []byte

	id0  id.ID
	id1  id.ID
	ttl  ttl.TTL
	body []byte
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

func (pack *PeerPacket) CreatePacketFromBytes(buf []byte) (Packet, error) {
	if buf == nil {
		return &PeerPacket{}, fmt.Errorf("Buffer is nil. Cannot create packet.")
	}
	var err error
	var newPacket *PeerPacket
	newPacket = new(PeerPacket)

	myTTL := new(ttl.PeerTTL)
	myID := new(id.PeerID)

	idLen := myID.GetLengthInBytes()
	ttlLen := myTTL.GetLengthInBytes()
	newPacket.id0, err = myID.CreateFromBytes(buf[0:idLen])
	if err != nil {
		return &PeerPacket{}, fmt.Errorf("Error parsing ID from buffer: %v", err)
	}
	newPacket.id1, err = myID.CreateFromBytes(buf[idLen : idLen*2])
	if err != nil {
		return &PeerPacket{}, fmt.Errorf("Error parsing ID from buffer: %v", err)
	}
	newPacket.ttl, err = myTTL.CreateFromBytes(buf[idLen*2 : idLen*2+ttlLen])
	if err != nil {
		return &PeerPacket{}, fmt.Errorf("Error parsing TTL from buffer: %v", err)
	}

	newPacket.body = buf[idLen*2+ttlLen:]

	return newPacket, nil
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
