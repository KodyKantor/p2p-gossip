//Receiver implements the udp Receiver interface.
package udp

import (
	"fmt"
	"net"

	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/udp/peer"
)

func init() {
	fmt.Println("Initialized receiver")
}

type Receiver struct {
	peer udp.Peer
}

//Receive receives packets from the network, turns it into a packet.Packet type,
// and then sends it through the provided channel.
func (r *Receiver) Receive(ch chan packet.Packet) error {
	listenStr := "localhost" + string(r.peer.GetPort())
	listenAddr, err := net.ResolveUDPAddr("udp", listenStr)
	if err != nil {
		return fmt.Errorf("Error resolving udp address %v: %v", listenStr, err)
	}
	recConn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		return fmt.Errorf("Error setting up listen connection: %v", err)
	}
	defer recConn.Close() //close the connection before the function returns

	buf := make([]byte, r.peer.GetPacketSize()) //make a buffer to hold the max packet size
	//TODO what do the other returns do?
	_, _, err = recConn.ReadFromUDP(buf)
	if err != nil {
		return fmt.Errorf("Error reading from udp: %v", err)
	}

	pack := new(packet.PeerPacket)
	newPacket, err := pack.CreatePacketFromBytes(buf)
	if err != nil {
		return fmt.Errorf("Error creating packet from udp buffer: %v", err)
	}

	ch <- newPacket
	fmt.Println("Sent packet through channel")

	fmt.Println("Exiting Receiver.")
	return nil
}
