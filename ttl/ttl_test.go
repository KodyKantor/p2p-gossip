package ttl

import (
	"testing"
)

//TestNew tests the CreateTTL() function.
func TestCreateTTL(t *testing.T) {
	var peerTTL *PeerTTL

	peerTTL = new(PeerTTL)

	val := 123
	myTTL, err := peerTTL.CreateTTL(val)
	if err != nil {
		t.Errorf("Error creating new ttl: %v", err)
	}
	if myTTL.GetTTL() != val {
		t.Errorf("TTL value is not correct: %v != %v", myTTL.GetTTL(), val)
	}

	val = -1
	myTTL, err = peerTTL.CreateTTL(val)
	if err == nil {
		t.Errorf("CreateTTL should have returned an error for a ttl val of %v", val)
	}
}

func TestSetTTL(t *testing.T) {
	var peerTTL *PeerTTL

	peerTTL = new(PeerTTL)

	val := 11
	err := peerTTL.SetTTL(val)
	if peerTTL.ttl != val {
		t.Errorf("TTL value is not correct: %v != %v", peerTTL.ttl, val)
	}
	if err != nil {
		t.Errorf("SetTTL threw an error: %v", err)
	}

	val = -12
	if err = peerTTL.SetTTL(val); err == nil {
		t.Errorf("SetTTL should have returned an error for value: %v", val)
	}
}

//TestGetBytes tests the GetBytes method.
func TestGetBytes(t *testing.T) {
	var peerTTL *PeerTTL

	peerTTL = new(PeerTTL)
	peerTTL.SetTTL(123)

	var buf []byte
	buf, err := peerTTL.GetBytes()
	if err != nil {
		t.Errorf("Error converting ttl to bytes: %v", err)
	}
	if len(buf) != 4 {
		t.Errorf("Len of byte slice must be 4 to be compatible with protocol: %v", len(buf))
	}
}

//TestGetTTL tests the GetTTL() function.
func TestGetTTL(t *testing.T) {
	var peerTTL *PeerTTL

	peerTTL = new(PeerTTL)

	val := 123
	if err := peerTTL.SetTTL(val); err != nil {
		t.Errorf("Error setting ttl value to %v: %v", val, err)
	}
}

func TestGetFromBytes(t *testing.T) {
	var peerTTL TTL

	peerTTL = new(PeerTTL)
	num := make([]byte, 4, 4) //allocate space and pointer

	num[0] = 1 // make the ttl = 1

	ret, err := peerTTL.CreateFromBytes(num)
	if err != nil {
		t.Errorf("Error creating TTL from byte slice: %v", err)
	}
	if ret.GetTTL() != 1 {
		t.Errorf("TTL value is incorrect: %v != %v", ret.GetTTL(), 1)
	}

}
