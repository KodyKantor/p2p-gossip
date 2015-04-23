//Receiver implements the udp Receiver interface.
package receiver

import (
	"fmt"
	"net"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/udp"
)

func init() {
	logrus.Debugln("Initialized receiver")
}

type Receiver struct {
	peer udp.Peer
}

func New(peer udp.Peer) *Receiver {
	//TODO check if peer is null
	newReceiver := new(Receiver)
	newReceiver.peer = peer
	return newReceiver
}

//Receive receives packets from the network, turns it into a packet.Packet type,
// and then sends it through the provided channel.
func (r *Receiver) Receive(ch chan packet.Packet) error {
	listenStr := "localhost:" + strconv.Itoa(r.peer.GetPort()) //convert port number to string
	listenAddr, err := net.ResolveUDPAddr("udp", listenStr)    //resolve the listen address
	if err != nil {
		return fmt.Errorf("Error resolving udp address %v: %v", listenStr, err)
	}
	logrus.Debugln("Resolved listen udp address.")

	recConn, err := net.ListenUDP("udp", listenAddr) //set up udp listener
	if err != nil {
		return fmt.Errorf("Error setting up listen connection: %v", err)
	}
	defer recConn.Close() //close the connection before the function returns
	logrus.Debugln("Got listen connection from listen address.")

	buf := make([]byte, r.peer.GetPacketSize()) //make a buffer to hold the max packet size
	//TODO what do the other returns do?
	_, _, err = recConn.ReadFromUDP(buf) //read from the udp connection (blocking)
	if err != nil {
		return fmt.Errorf("Error reading from udp: %v", err)
	}
	logrus.Debugln("Read packet from udp.")

	//Create a packet from the new buffer.
	pack := new(packet.PeerPacket)
	newPacket, err := pack.CreatePacketFromBytes(buf)
	if err != nil {
		return fmt.Errorf("Error creating packet from udp buffer: %v", err)
	}
	logrus.Debugln("Created packet type from packet buffer.")

	ch <- newPacket //send the formatted packet through the channel
	logrus.Debugln("Sent packet through channel")

	logrus.Debugln("Exiting Receiver.")
	return nil
}
