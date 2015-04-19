package id

import (
	"github.com/Sirupsen/logrus"
)

const (
	DEFAULT_SIZE int = 32
)

var log = logrus.New()

func init() {
	log.Println("Initialized id")
}

type ID interface {
	ServeIDs(chan ID) // sends IDs through the provided channel
	Equals(ID) bool   // tells whether or not two ids are equal
	GetBytes() []byte // returns a byte-slice representation of the ID
	GetLength() int   // returns the number of bytes needed for the ID
	SetLength(int)    // sets the number of bytes an ID uses
	CreateFromBytes([]byte) (ID, error)
	GetZeroID() (ID, error)

	createID() (ID, error)
}
