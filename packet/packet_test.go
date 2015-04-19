package packet

import (
	"testing"
)

//TestBufferable implements the Bufferable interface,
//and is used to test the PeerPacket type.
type TestBufferable struct {
	GetBytesCalled bool

	Buffer []byte //buffer for testing
}

func (t *TestBufferable) GetBytes() []byte {
	t.GetBytesCalled = true

	return t.Buffer
}

func (t *TestBufferable) Equals(other []byte) bool {
	if len(t.Buffer) != len(other) {
		return false
	}
	for i := 0; i < len(t.Buffer); i++ {
		if t.Buffer[i] != other[i] {
			return false
		}
	}

	return true
}

//AggregateEquals tests to see if one mega-buffer is equal
//to the sum of the following buffers.
func AggregateEquals(megabuf []byte, bufs ...[]byte) bool {
	megabuflen := len(megabuf)

	aggregatelen := 0
	for _, buf := range bufs {
		aggregatelen = aggregatelen + len(buf)
	}
	if megabuflen != aggregatelen {
		return false
	}

	k := 0
	for i := 0; i < len(bufs); i++ {
		for j := 0; j < len(bufs[i]); j++ {
			//			fmt.Printf("megabuf[%v] = %v, bufs[%v][%v] = %v\n", k, megabuf[k], i, j, bufs[i][j])
			if megabuf[k] != bufs[i][j] {
				return false
			}
			k++
		}
	}

	return true
}

func TestCreatePacket(t *testing.T) {
	myBuffer := TestBufferable{false, nil}

	pack := new(PeerPacket)
	if _, err := pack.CreatePacket(&myBuffer); err == nil {
		t.Error("CreatePacket should return an error for a nil buffer.")
	}
	if myBuffer.GetBytesCalled == false {
		t.Error("GetBytes should have been called, but was not.")
	}
}

func TestGetBuffer(t *testing.T) {

	//test a single buffer
	myBuffer := TestBufferable{false, nil}
	myBuffer.Buffer = []byte{0, 1, 2, 3, 4}

	pack := new(PeerPacket)
	newPacket, err := pack.CreatePacket(&myBuffer)

	if err != nil {
		t.Errorf("CreatePacket returned an error: %v", err)
	}
	if myBuffer.GetBytesCalled == false {
		t.Error("GetBytes should have been called, but was not.")
	}

	buf := newPacket.GetBuffer()
	if buf == nil {
		t.Error("GetBuffer returned a nil buffer")
	}
	if !myBuffer.Equals(buf) {
		t.Error("Buffers are not equal. %v != %v", myBuffer.Buffer, buf)
	}

	//test two buffers
	buf1 := TestBufferable{false, nil}
	buf2 := TestBufferable{false, nil}

	buf1.Buffer = []byte{0, 1, 2, 3, 4}
	buf2.Buffer = []byte{5, 6, 7, 8, 9}

	pack = new(PeerPacket)
	newPacket, err = pack.CreatePacket(&buf1, &buf2)

	if err != nil {
		t.Errorf("CreatePacket returned an error: %v", err)
	}
	if buf1.GetBytesCalled == false {
		t.Error("GetBytes should have been called, but was not.")
	}
	if buf2.GetBytesCalled == false {
		t.Error("GetBytes should have been called, but was not.")
	}

	buf = newPacket.GetBuffer()
	if buf == nil {
		t.Error("GetBuffer returned a nil buffer")
	}

	if !AggregateEquals(buf, buf1.Buffer, buf2.Buffer) {
		t.Errorf("Aggregation of buffers is not equal: %v != %v + %v", buf, buf1.Buffer, buf2.Buffer)
	}

}
