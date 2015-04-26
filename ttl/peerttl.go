package ttl

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/Sirupsen/logrus"
)

const (
	//DecrementValue is the default value to use when decrementing TTLs.
	DecrementValue = 1
	//DefaultLength is the default length in bytes of generated TTLs.
	DefaultLength = 4
)

//PeerTTL is an implementation of the TTL interface.
type PeerTTL struct {
	ttl    int //time to live
	length int
}

//NewTTL returns a pointer to a new TTL type.
func NewTTL() *PeerTTL {
	return &PeerTTL{0, DefaultLength}
}

//SetTTL sets the ttl attribute.
func (t *PeerTTL) SetTTL(ttl int) error {
	if ttl < 0 {
		return fmt.Errorf("Invalid time to live: %v", ttl)
	}
	t.ttl = ttl
	return nil
}

//GetTTL returns the ttl attribute.
func (t *PeerTTL) GetTTL() int {
	return t.ttl
}

//GetBytes converts the TTL to a byte slice. Returns an error if the byte slice isn't four bytes.
func (t *PeerTTL) GetBytes() []byte {
	buf := make([]byte, 4) //ints are 4 bytes
	binary.PutVarint(buf, int64(t.ttl))
	return buf
}

//CreateTTL takes an integer, and converts it into a TTL.
func (t *PeerTTL) CreateTTL(time int) (TTL, error) {
	if time < 0 {
		return &PeerTTL{}, fmt.Errorf("Invalid time to live: %v", time)
	}
	return &PeerTTL{time, DefaultLength}, nil
}

// CreateFromBytes takes a byte slice and turns it into a TTL.
func (t *PeerTTL) CreateFromBytes(time []byte) (TTL, error) {
	logrus.Debugln("Entered TTL CreateFromBytes.")
	if time == nil {
		return &PeerTTL{}, fmt.Errorf("TTL byte slice is nil")
	}
	if len(time) == 0 {
		return &PeerTTL{}, fmt.Errorf("TTL byte slice is zero-length")
	}

	logrus.Debugln("Creating new TTL from the buffer:", time)
	var ret int32 //to hold the decoded value
	buf := bytes.NewBuffer(time)
	err := binary.Read(buf, binary.LittleEndian, &ret)
	//	decoded, err := binary.ReadVarint(buf)
	if err != nil {
		return &PeerTTL{}, fmt.Errorf("Error deocding ttl: %v", err)
	}
	return &PeerTTL{int(ret), DefaultLength}, nil
}

//DecrementTTL decrements the TTL by the constant value defined in the package.
func (t *PeerTTL) DecrementTTL() {
	t.ttl = t.ttl - DecrementValue
}

//GetLengthInBytes returns the length of the generated TTLs (in bytes).
func (t *PeerTTL) GetLengthInBytes() int {
	return t.length
}
