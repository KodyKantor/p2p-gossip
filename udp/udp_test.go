package udp

import (
	"net"
	"testing"
)

const (
	senderAddr   = "localhost:8090"
	receiverAddr = "localhost:8080"
)

func TestUDPConnectionless(t *testing.T) {
	t.Log("Beginning test")

	la, err := net.ResolveUDPAddr("udp", receiverAddr)
	if err != nil {
		t.Fatalf("Error resolving receiver address: %v", err)
	}

	t.Log("Getting sender")
	sendConn, err := net.ListenPacket("udp", senderAddr) //sender's connection
	if err != nil {
		t.Fatalf("Error getting sender's connection established: %v", err)
	}
	defer sendConn.Close()

	t.Log("Getting listener")
	recConn, err := net.ListenUDP("udp", la)
	if err != nil {
		t.Fatalf("Error getting receiver's connection established: %v", err)
	}
	defer recConn.Close()

	t.Log("Starting receiver")
	recChan := make(chan int, 1)
	go func(ch chan int) {
		buf := make([]byte, 1024)
		_, _, err := recConn.ReadFromUDP(buf)
		if err != nil {
			t.Fatalf("Error reading from UDP: %v", err)
		}
		t.Log("Buffer said: ", string(buf))
		recChan <- 0
	}(recChan)

	// send a string
	_, err = sendConn.(*net.UDPConn).WriteToUDP([]byte("Hello receiver!"), la)
	if err != nil {
		t.Fatalf("Error writing to UDP: %v", err)
	}
	t.Log("Wrote to UDP connection")

	//exit
	<-recChan

	t.Log("Exiting.")

}
