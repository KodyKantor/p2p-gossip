//Package id provides an interface that defines a random N-byte id generator,
//and an implementation of the ID interface.
package id

import (
	"github.com/Sirupsen/logrus"
)

const (
	//DefaultSize is the default size of generated IDs.
	DefaultSize int = 32
)

var log = logrus.New()

func init() {
	log.Debugln("Initialized id")
}

//ID interface provides functionality for creating bufferizable IDs to send
// in packets.
type ID interface {
	ServeIDs(chan ID)      // sends IDs through the provided channel
	Equals(ID) bool        // tells whether or not two ids are equal
	GetBytes() []byte      // returns a byte-slice representation of the ID
	GetLengthInBytes() int // returns the number of bytes needed for the ID
	SetLength(int)         // sets the number of bytes an ID uses
	CreateFromBytes([]byte) (ID, error)
	GetZeroID() (ID, error)

	createID() (ID, error)
}
