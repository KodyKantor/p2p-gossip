package packet

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/id"
	"github.com/kodykantor/p2p-gossip/ttl"
)

//PeerPacket implements the Packet interface.
type PeerPacket struct {
	payload       []Bufferable //payload is a slice of bufferable things
	bufferization []byte

	Id0  id.ID
	Id1  id.ID
	Ttl  ttl.TTL
	body []byte
}

func NewPacket() *PeerPacket {
	return &PeerPacket{}
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
	logrus.Debugln("Entered CreatePacketFromBytes with buffer:", buf)
	if buf == nil {
		return &PeerPacket{}, fmt.Errorf("Buffer is nil. Cannot create packet.")
	}
	var err error

	newPacket := NewPacket()
	myTTL := ttl.NewTTL()
	myID := id.NewID()

	idLen := myID.GetLengthInBytes()
	ttlLen := myTTL.GetLengthInBytes()

	tmpBuf := buf[0:idLen]
	logrus.Debugln("Creating Id0 from bytes:", tmpBuf)
	newPacket.Id0, err = myID.CreateFromBytes(tmpBuf)
	if err != nil {
		return &PeerPacket{}, fmt.Errorf("Error parsing ID from buffer: %v", err)
	}

	logrus.Debugln("Creating Id1 from bytes.")
	newPacket.Id1, err = myID.CreateFromBytes(buf[idLen : idLen*2])
	if err != nil {
		return &PeerPacket{}, fmt.Errorf("Error parsing ID from buffer: %v", err)
	}

	logrus.Debugln("Creating ttl from bytes.")
	newPacket.Ttl, err = myTTL.CreateFromBytes(buf[idLen*2 : idLen*2+ttlLen])
	if err != nil {
		return &PeerPacket{}, fmt.Errorf("Error parsing TTL from buffer: %v", err)
	}

	logrus.Debugln("Setting body to the rest of the packet.")
	newPacket.body = buf[idLen*2+ttlLen:]

	logrus.Debugln("Returning new packet.")
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

func (pack *PeerPacket) GetBufferization() []byte {
	return pack.bufferization
}
