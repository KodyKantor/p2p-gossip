package id

import (
	"testing"
)

//587 - 612

//TestNewID tests that the New() function
//works properly.
func TestNewID(t *testing.T) {
	var newID ID
	newID = New(32)

	if newID.GetID() == nil {
		t.Error("Randomly generated ID is nil")
	}

	t.Log("Generated ID is", newID.randomID)
}

//TestGetID tests the GetID() function.
func TestGetID(t *testing.T) {
	var newID ID

	newID = New(32)
	toTest := newID.GetID()

	for i := 0; i < len(newID.randomID); i++ {
		if newID.randomID[i] != toTest[i] {
			t.Error("IDs do not match")
		}
	}
}

//TestEquals tests the Equals() function.
func TestEquals(t *testing.T) {
	var newID ID
	var otherID ID

	newID = New(32)
	otherID = New(32)

	if newID.Equals(newID) == false {
		t.Error("The ID did not equal itself")
	}

	if newID.Equals(otherID) == true {
		t.Error("Two different IDs should not be equal")
	}
}

//TestIDSize tests the GetIDSize() and SetIDSize() functions.
func TestIDSize(t *testing.T) {
	var newID ID

	size := 32
	newID = New(size)

	if newID.GetIDSize() != size {
		t.Error("Sizes do not match")
	}

	size = 64
	newID.SetIDSize(size)
	if newID.GetIDSize() != size {
		t.Error("Sizes do not match")
	}

}

//TestServeIDs will start the ID server, and read
//a series of IDs from the channel.
func TestServeIDs(t *testing.T) {
	var generatorID ID
	ch := make(chan ID, 1)

	generatorID = New(32)
	go generatorID.ServeIDs(ch) //start the server

	for i := 0; i < 50; i++ {
		newID := <-ch
		if newID.GetID() == nil {
			t.Error("Randomly generated ID is nil")
		}
	}
}

//TestGetZeroID tests the GetZeroID function.
func TestGetZeroID(t *testing.T) {
	var newID ID
	var err error

	for size := 32; size < 64; size++ {
		newID, err = GetZeroID(size)
		if err != nil {
			t.Errorf("Error should be nil: %v", err)
		}
		if len(newID.randomID) != size {
			t.Errorf("Randomly generated ID is not the proper size: %v != %v", len(newID.randomID), size)
		}
	}

	for size := -64; size < 1; size++ {
		if _, err = GetZeroID(size); err == nil {
			t.Errorf("GetZeroID should return an error for a size of %v", size)
		}
	}
}
