package ttl

import (
	"testing"
)

//TestNew tests the New() function.
func TestNew(t *testing.T) {
	var newTTL TTL

	val := 123
	newTTL = New(val)
	if newTTL.ttl != val {
		t.Errorf("TTL value is not correct: %v != %v", newTTL.ttl, val)
	}

	val = -1
	newTTL = New(val)
	if newTTL.ttl != DEFAULT_TTL {
		t.Errorf("TTL value should be the default value: %v != %v", newTTL.ttl, val)
	}
}

func TestSetTTL(t *testing.T) {
	var newTTL TTL

	val := 123
	newTTL = New(val)

	val = 11
	err := newTTL.SetTTL(val)
	if newTTL.ttl != val {
		t.Errorf("TTL value is not correct: %v != %v", newTTL.ttl, val)
	}
	if err != nil {
		t.Errorf("SetTTL threw an error: %v", err)
	}

	val = -12
	if err = newTTL.SetTTL(val); err == nil {
		t.Errorf("SetTTL should have returned an error for value: %v", val)
	}
}
