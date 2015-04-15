package id

import (
	"crypto/rand"
	"github.com/Sirupsen/logrus"
)

const (
	DEFAULT_SIZE int = 32
)

var log = logrus.New()
var zeroID = make([]byte, DEFAULT_SIZE) //can be edited later

func init() {
	log.Println("Initialized id")
}

type ID struct {
	randomID []byte
	size     int
}

//New returns an ID struct
func New(numberOfBytes int) ID {
	id := generateID(numberOfBytes)
	return ID{id, numberOfBytes}
}

func (id *ID) ServeIDs(c chan ID) {
	for true {
		log.Println("Sending new ID")
		c <- New(id.size)
	}
}

//Equals compares two IDs.
func (id *ID) Equals(other ID) bool {
	//if the sizes aren't the same, don't bother iterating
	//through the byte slices.
	if id.size != other.size {
		return false
	}
	for i := 0; i < id.size; i++ {
		if id.randomID[i] != other.randomID[i] {
			return false
		}
	}

	return true
}

//GenerateID creates a byte slice to contain
//a bunch of cryptographically random bytes.
//The numberOfBytes parameter is the size of the
//randomly generated slice.
func generateID(numberOfBytes int) []byte {
	slice := make([]byte, numberOfBytes)
	_, err := rand.Read(slice)
	if err != nil {
		log.Errorln("Error creating random id")
		return nil
	}

	sum := 0
	for i := 0; i < numberOfBytes; i++ {
		sum = sum + int(slice[i])
	}
	if sum != 0 {
		return slice
	}
	//if the id generated is zero, we need to try again
	return generateID(numberOfBytes)
}

//GetZeroID returns a new ID structure with a totally
//zeroed random ID.
func GetZeroID(numberOfBytes int) ID {
	return ID{zeroID, numberOfBytes}
}

//GetID copies the ID before returning it.
func (id *ID) GetID() []byte {
	return id.randomID //this is actually safe in Go
}

//GetIDSize returns the private size variable in the struct
func (id *ID) GetIDSize() int {
	return id.size
}

//SetIDSize sets the size in bytes of the
//future-generated IDs.
func (id *ID) SetIDSize(size int) {
	if size < 0 {
		size = DEFAULT_SIZE
	}
	id.size = size
}
