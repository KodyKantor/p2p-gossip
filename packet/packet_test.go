package packet

import (
	"testing"

	"github.com/kodykantor/p2p-gossip/id"
	"github.com/kodykantor/p2p-gossip/ttl"
)

//TestBufferable implements the Bufferable interface,
//and is used to test the PeerPacket type.
type TestBufferable struct {
	GetBytesCalled bool   //tells if the GetBytes() func was called
	Buffer         []byte //buffer for testing
}

func (t *TestBufferable) GetBytes() []byte {
	t.GetBytesCalled = true

	return t.Buffer
}

func TestCreatePacketFromBytes(t *testing.T) {
	myID := id.NewID()
	myTTL := ttl.NewTTL()
	idLen := myID.GetLengthInBytes()
	ttlLen := myTTL.GetLengthInBytes()
	bodyLen := 20

	idbuf0 := make([]byte, idLen)
	idbuf1 := make([]byte, idLen)
	ttlbuf := make([]byte, ttlLen)
	bodybuf := make([]byte, bodyLen) //20 bytes for the body

	for i := 0; i < idLen; i++ {
		idbuf0[i] = byte(i)
	}
	t.Logf("idbuf0 is %v", idbuf0)
	for i := 1; i < idLen; i++ {
		idbuf1[idLen-i] = byte(i)
	}
	t.Logf("idbuf1 is %v", idbuf1)
	for i := 0; i < ttlLen; i++ {
		ttlbuf[i] = byte(i)
	}
	t.Logf("ttlbuf is %v", ttlbuf)
	for i := 0; i < bodyLen; i++ {
		bodybuf[i] = byte(i)
	}
	t.Logf("bodybuf is %v", bodybuf)

	//make a mega buffer (a packet)
	x := append(idbuf0, idbuf1...)
	x = append(x, ttlbuf...)
	x = append(x, bodybuf...)

	testPacket := new(PeerPacket)
	pack, err := testPacket.CreatePacketFromBytes(x)
	if err != nil {
		t.Errorf("Error creating packet from megabuffer: %v", err)
	}

	id0 := pack.(*PeerPacket).ID0
	id1 := pack.(*PeerPacket).ID1
	ttl := pack.(*PeerPacket).TTL
	body := pack.(*PeerPacket).Body

	if id0 == nil || id1 == nil || ttl == nil {
		t.Error("IDs or TTL is nil")
	}

	slice := id0.GetBytes()
	for ind, val := range slice {
		if val != idbuf0[ind] {
			t.Errorf("Expected %v, got %v", idbuf0[ind], val)
		}
	}

	slice = id1.GetBytes()
	for ind, val := range slice {
		if val != idbuf1[ind] {
			t.Errorf("Expected %v, got %v", idbuf1[ind], val)
		}
	}

	//Since ttls are an int represented by a byte slice,
	// we have to massage the data we sent in a little bit.
	slice = ttl.GetBytes()
	newTTL, err := myTTL.CreateFromBytes(ttlbuf)
	if err != nil {
		t.Errorf("Error creating ttl from bytes. Check the ttl functions. %v", err)
	}
	ttlBufAsBytes := newTTL.GetBytes()

	t.Logf("Received %v", slice)
	t.Logf("Comparing to %v", ttlbuf)
	for ind, val := range slice {
		if val != ttlBufAsBytes[ind] {
			t.Errorf("Expected %v, got %v", ttlBufAsBytes[ind], val)
		}
	}

	t.Logf("Received %v", body)
	t.Logf("Comparing to %v", bodybuf)
	for ind, val := range body {
		if val != bodybuf[ind] {
			t.Errorf("Expected %v, got %v", bodybuf[ind], val)
		}
	}

	_, err = testPacket.CreatePacketFromBytes(nil)
	if err == nil {
		t.Error("CreatePacketFromBytes should return an error for nil slice")
	}

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

	buf := newPacket.GetBufferization()
	if buf == nil {
		t.Error("GetBuffer returned a nil buffer")
	}
	if !myBuffer.Equals(buf) {
		t.Errorf("Buffers are not equal. %v != %v", myBuffer.Buffer, buf)
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

	buf = newPacket.GetBufferization()
	if buf == nil {
		t.Error("GetBuffer returned a nil buffer")
	}

	if !AggregateEquals(buf, buf1.Buffer, buf2.Buffer) {
		t.Errorf("Aggregation of buffers is not equal: %v != %v + %v", buf, buf1.Buffer, buf2.Buffer)
	}

}

func TestBufferizableString(t *testing.T) {
	s := BufferizableString{Str: "hello there!"}
	slice := s.GetBytes()
	if slice == nil {
		t.Error("Buffer is nil")
	}
}
