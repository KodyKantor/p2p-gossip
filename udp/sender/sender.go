package sender

import (
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/udp/peer"
)

//Sender implements the Sender interface.
type Sender struct {
	peer *peer.Peer
}

func init() {
	logrus.Debugln("Initialized Sender package.")
}

//New returns a pointer to a Sender type. Takes an initialized peer pointer.
func New(peer *peer.Peer) *Sender {
	//TODO check if peer is nil
	newSender := new(Sender)
	newSender.peer = peer
	return newSender
}

const (
	senderAddr   = "localhost:12349"
	receiverAddr = "localhost:12345"
)

//Send sends packets through UDP to the Sender's defined peer partners.
//If an error occurs, an error is returned. Packets that need to be sent are
//passed in through a packet channel.
func (s *Sender) Send(ch chan *packet.PeerPacket) error {
	logrus.Debugln("Starting to send packet...")
	//	senderAddr := "localhost:" + strconv.Itoa(s.peer.GetPort()) //TODO change this
	//	receiverAddr := "localhost:" + strconv.Itoa(s.peer.GetPort())

	var packetToSend *packet.PeerPacket

	logrus.Printf("Client's listen address should be: %v", receiverAddr)
	la, err := net.ResolveUDPAddr("udp", receiverAddr)
	if err != nil {
		return fmt.Errorf("Error resolving the listener's address: %v", err)
	}

	logrus.Debugln("Starting packet listener for sender")
	sendConn, err := net.ListenPacket("udp", senderAddr)
	if err != nil {
		return fmt.Errorf("Error getting send connection: %v", err)
	}
	defer sendConn.Close()

	for true {
		logrus.Debugln("Reading from channel to send.")
		packetToSend = <-ch //read a packet to send from the channel
		logrus.Debugln("Received packet from channel to send.")
		buf := packetToSend.GetBufferization()

		count, err := sendConn.(*net.UDPConn).WriteToUDP(buf, la)
		if err != nil {
			return fmt.Errorf("Error writing packet to UDP: %v", err)
		}
		logrus.Debugln("Sent ", count, "bytes.")
		logrus.Debugln("Sender sent packet!")
	}
	return nil
}
