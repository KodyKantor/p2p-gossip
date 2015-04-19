package ttl

import (
	"encoding/binary"
	"fmt"
)

const DECREMENT_VALUE = 1

type PeerTTL struct {
	ttl int //time to live
}

//SetPeerTTL sets the ttl attribute.
func (t *PeerTTL) SetTTL(ttl int) error {
	if ttl < 0 {
		return fmt.Errorf("Invalid time to live: %v", ttl)
	}
	t.ttl = ttl
	return nil
}

//GetPeerTTL returns the ttl attribute.
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
	return &PeerTTL{time}, nil
}

// CreateFromBytes takes a byte slice and turns it into a TTL.
func (t *PeerTTL) CreateFromBytes(time []byte) (TTL, error) {
	log.Println("size of slice is ", len(time))
	decoded, count := binary.Varint(time)
	log.Println("Decoded the value", decoded)
	res := int(decoded) // cast as 32-bit integer
	if count != 4 {
		return &PeerTTL{}, fmt.Errorf("Improper number of bytes read from buffer: %v", count)
	}
	return &PeerTTL{res}, nil
}

//DecrementTTL decrements the TTL by the constant value defined in the package.
func (t *PeerTTL) DecrementTTL() {
	t.ttl = t.ttl - DECREMENT_VALUE
}
