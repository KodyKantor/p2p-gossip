package udp

import (
	"fmt"
	"net"
	"testing"
	"time"
)

const (
	senderAddr   = "localhost:8090"
	receiverAddr = "localhost:8080"
)

func TestUDPConnectionless(t *testing.T) {
	fmt.Println("Beginning test")

	recChan := make(chan int, 1)
	la, err := net.ResolveUDPAddr("udp", receiverAddr)
	if err != nil {
		t.Fatalf("Error resolving receiver address: %v", err)
	}

	go runReceiver(recChan, la, t)
	time.Sleep(time.Second * 1)
	runSender(la, t)

	//exit
	<-recChan

	fmt.Println("Exiting.")

}

func runReceiver(recChan chan int, la *net.UDPAddr, t *testing.T) {
	fmt.Println("Getting listener")

	recConn, err := net.ListenUDP("udp", la)
	if err != nil {
		t.Fatalf("Error getting receiver's connection established: %v", err)
	}
	defer recConn.Close()

	fmt.Println("Starting receiver")

	buf := make([]byte, 1024)
	_, _, err = recConn.ReadFromUDP(buf)
	if err != nil {
		t.Fatalf("Error reading from UDP: %v", err)
	}
	fmt.Println("Buffer said: ", string(buf))
	recChan <- 0
}

//Runs the sender
func runSender(la *net.UDPAddr, t *testing.T) {
	fmt.Println("Getting sender")
	sendConn, err := net.ListenPacket("udp", senderAddr) //sender's connection: localhost:8090
	if err != nil {
		t.Fatalf("Error getting sender's connection established: %v", err)
	}
	defer sendConn.Close()

	// send a string
	_, err = sendConn.(*net.UDPConn).WriteToUDP([]byte("Hello receiver!"), la) //send thru sender's conn, to listen addr: localhost:8080
	if err != nil {
		t.Fatalf("Error writing to UDP: %v", err)
	}
	fmt.Println("Wrote to UDP connection")
}
