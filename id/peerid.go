package id

import (
	"crypto/rand"
	"fmt"

	"github.com/Sirupsen/logrus"
)

//PeerID is an implementation of the ID interface.
type PeerID struct {
	randomID []byte //byte slice to hold ID
	size     int    // size in bytes of ID
}

//NewID returns a pointer to a PeerID type.
func NewID() *PeerID {
	return &PeerID{nil, DefaultSize}
}

//createID creates a byte slice to contain
//a bunch of cryptographically random bytes.
//The numberOfBytes parameter is the size of the
//randomly generated slice.
func (id *PeerID) createID() (ID, error) {
	log.Debugln("Creating new ID")
	if id.size < 1 {
		return &PeerID{}, fmt.Errorf("Invalid ID size: %v", id.size)
	}

	slice := make([]byte, id.size)
	_, err := rand.Read(slice)
	if err != nil {
		return &PeerID{}, fmt.Errorf("Error creating random ID: %v", err)
	}

	sum := 0
	for i := 0; i < id.size; i++ {
		sum = sum + int(slice[i])
	}
	if sum != 0 { //we don't have a zero id
		return &PeerID{slice, id.size}, nil
	}

	return id.createID() //try again, because we got a zero id
}

//CreateFromBytes creates an ID from the provided byte slice.
func (id *PeerID) CreateFromBytes(slice []byte) (ID, error) {
	if slice == nil {
		return &PeerID{}, fmt.Errorf("Byte slice is nil. Cannot create ID")
	}

	logrus.Debugln("Creating id from the bytes:", slice)

	return &PeerID{slice, len(slice)}, nil
}

//ServeIDs returns ID structs through the provided channel.
func (id *PeerID) ServeIDs(c chan ID) {
	for true {
		log.Debug("Sending new ID")
		id, err := id.createID()
		if err != nil {
			log.Error("Error creating new id: ", err)
		}
		c <- id
	}
}

//Equals compares two IDs.
func (id *PeerID) Equals(other ID) bool {
	//if the sizes aren't the same, don't bother iterating
	//through the byte slices.
	if id.GetLengthInBytes() != other.GetLengthInBytes() {
		return false
	}

	mybytes := id.GetBytes()
	otherbytes := other.GetBytes()
	for i := 0; i < id.GetLengthInBytes(); i++ {
		if mybytes[i] != otherbytes[i] {
			return false
		}
	}

	return true
}

//GetBytes copies the ID before returning it.
func (id *PeerID) GetBytes() []byte {
	return id.randomID //this is actually safe in Go
}

//GetLengthInBytes returns the private size variable in the struct
func (id *PeerID) GetLengthInBytes() int {
	return id.size
}

//SetLength sets the size in bytes of the
//future-generated IDs.
func (id *PeerID) SetLength(size int) {
	if size < 1 {
		size = DefaultSize
	}
	id.size = size
}

//GetZeroID returns an ID type with an all-zero id.
func (id *PeerID) GetZeroID() (ID, error) {
	if id.size < 1 {
		return &PeerID{}, fmt.Errorf("Invalid id size: %v", id.size)
	}

	zeroID := make([]byte, id.size)
	return id.CreateFromBytes(zeroID)
}
