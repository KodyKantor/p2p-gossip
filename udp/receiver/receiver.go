//Package receiver implements the udp Receiver interface.
package receiver

import (
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/udp/peer"
)

func init() {
	logrus.Debugln("Initialized receiver")
}

//Receiver retains information about the running peer.
type Receiver struct {
	peer *peer.Peer
}

//New returns a pointer to a Receiver.
func New(peer *peer.Peer) *Receiver {
	//TODO check if peer is null
	newReceiver := new(Receiver)
	newReceiver.peer = peer
	return newReceiver
}

const (
	senderAddr   = "localhost:12349"
	receiverAddr = "localhost:12345"
	packetSize   = 512
)

//Receive receives packets from the network, turns it into a packet.Packet type,
// and then sends it through the provided channel.
func (r *Receiver) Receive(ch chan *packet.PeerPacket) error {
	//	senderAddr := "localhost:" + strconv.Itoa(r.peer.GetPort()) //convert port number to string
	listenAddr, err := net.ResolveUDPAddr("udp", receiverAddr) //resolve the listen address
	logrus.Printf("Client's actual listen address is: %v", receiverAddr)
	if err != nil {
		return fmt.Errorf("Error resolving udp address %v: %v", receiverAddr, err)
	}
	logrus.Debugln("Resolved listen udp address.")

	recConn, err := net.ListenUDP("udp", listenAddr) //set up udp listener
	if err != nil {
		return fmt.Errorf("Error setting up listen connection: %v", err)
	}
	defer recConn.Close() //close the connection before the function returns
	logrus.Debugln("Got listen connection from listen address.")

	logrus.Debugln("Peer packet size is ", packetSize)

	buf := make([]byte, packetSize) //make a buffer to hold the max packet size
	for true {                      // listen for packets forever
		count, _, err := recConn.ReadFromUDP(buf) //read from the udp connection (blocking)
		if err != nil {
			return fmt.Errorf("Error reading from udp: %v", err)
		}
		logrus.Debugln("Read packet from udp.")
		logrus.Debugln("Read ", count, "bytes from the connection.")
		logrus.Debugln("Packet read from buffer is:\n", buf)

		//Create a packet from the new buffer.
		pack := new(packet.PeerPacket)

		logrus.Debugln("Creating packet type from packet buffer.")
		newPacket, err := pack.CreatePacketFromBytes(buf)
		if err != nil {
			return fmt.Errorf("Error creating packet from udp buffer: %v", err)
		}
		logrus.Debugln("Created packet type from packet buffer.")

		ch <- newPacket.(*packet.PeerPacket) //send the formatted packet through the channel
		logrus.Debugln("Sent packet through channel")
	}
	logrus.Debugln("Exiting Receiver.")
	return nil
}
