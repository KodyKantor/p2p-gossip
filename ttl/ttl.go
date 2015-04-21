//Package ttl provides a TTL interface that defines a
//time-to-live that can be packetized. There is also an implementation
//of the TTL interface.
package ttl

import (
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

const (
	DEFAULT_TTL = 30
)

func init() {
	log.Println("Initialized ttl")
}

//TTL implements the Bufferable interface.
type TTL interface {
	GetBytes() []byte                    //returns raw ttl as a byte slice
	GetTTL() int                         //returns the TTL as an integer
	CreateTTL(int) (TTL, error)          //creates a new TTL with set length
	SetTTL(int) error                    //changes TTL
	CreateFromBytes([]byte) (TTL, error) //creates TTL from raw byte slice
	DecrementTTL()                       //decrements TTL
	GetLengthInBytes() int
}
