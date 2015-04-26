package id

import (
	"testing"
)

//587 - 612

//TestNewID tests that the New() function
//works properly.
func TestNewID(t *testing.T) {
	var newID ID

	newID = new(PeerID)
	newID.SetLength(32)

	myid, err := newID.createID()
	if err != nil {
		t.Errorf("createID returned an error: %v", err)
	}
	if myid.GetBytes() == nil {
		t.Error("Randomly generated ID is nil")
	}

	t.Log("Generated ID is", myid.GetBytes())
}

//TestEquals tests the Equals() function.
func TestEquals(t *testing.T) {
	var newID ID

	newID = new(PeerID)
	newID.SetLength(32)

	id0, err := newID.createID()
	if err != nil {
		t.Errorf("createID returned an error: %v", err)
	}

	id1, err := newID.createID()
	if err != nil {
		t.Errorf("createID returned an error: %v", err)
	}

	if id0.Equals(id0) == false {
		t.Error("The ID did not equal itself")
	}

	if id0.Equals(id1) == true {
		t.Error("Two different IDs should not be equal")
	}
}

//TestIDSize tests the GetLengthInBytes() and SetIDSize() functions.
func TestIDSize(t *testing.T) {
	var newID ID

	size := 32
	newID = new(PeerID)
	newID.SetLength(size)

	if newID.GetLengthInBytes() != size {
		t.Error("Sizes do not match")
	}

	size = 64
	newID.SetLength(size)
	if newID.GetLengthInBytes() != size {
		t.Error("Sizes do not match")
	}

	for size := -64; size < 1; size++ {
		newID.SetLength(size)
		if newID.GetLengthInBytes() != DefaultSize {
			t.Errorf("SetID should return an error for a size of %v", size)
		}
	}

}

//TestServeIDs will start the ID server, and read
//a series of IDs from the channel.
func TestServeIDs(t *testing.T) {
	var generatorID ID
	ch := make(chan ID, 1)

	generatorID = new(PeerID)
	generatorID.SetLength(32)
	go generatorID.ServeIDs(ch) //start the server

	for i := 0; i < 50; i++ {
		newID := <-ch
		if newID.GetBytes() == nil {
			t.Error("Randomly generated ID is nil")
		}
	}
}

//TestGetZeroID tests the GetZeroID function.
func TestGetZeroID(t *testing.T) {
	var newID ID

	newID = new(PeerID)
	for size := 32; size < 64; size++ {
		newID.SetLength(size)
		zeroid, err := newID.GetZeroID()
		if err != nil {
			t.Errorf("Error should be nil: %v", err)
		}
		if len(zeroid.GetBytes()) != size {
			t.Errorf("Randomly generated ID is not the proper size: %v != %v", len(zeroid.GetBytes()), size)
		}
	}

}
