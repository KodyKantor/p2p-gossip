package sender

import (
	"fmt"
	"net"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/udp"
)

type Sender struct {
	peer udp.Peer
}

func init() {
}

func New(peer udp.Peer) *Sender {
	//TODO check if peer is nil
	newSender := new(Sender)
	newSender.peer = peer
	return newSender
}

func (s *Sender) Send(ch chan packet.Packet) error {
	logrus.Debugln("Starting to send packet...")
	sendStr := "localhost:" + strconv.Itoa(s.peer.GetPort()+10) //TODO change this
	clientListenAddr := "localhost:" + strconv.Itoa(s.peer.GetPort())

	logrus.Printf("Resolving udp address for client: %v", clientListenAddr)
	la, err := net.ResolveUDPAddr("udp", clientListenAddr)
	if err != nil {
		return fmt.Errorf("Error resolving the listener's address: %v", err)
	}

	logrus.Debugln("Starting packet listener for sender")
	sendConn, err := net.ListenPacket("udp", sendStr)
	if err != nil {
		return fmt.Errorf("Error getting send connection: %v", err)
	}
	defer sendConn.Close()

	logrus.Debugln("Reading from channel to send.")
	pack := <-ch //read a packet from the channel
	logrus.Debugln("Received packet from channel to senD.")
	buf := pack.GetBuffer()
	_, err = sendConn.(*net.UDPConn).WriteToUDP(buf, la)
	if err != nil {
		return fmt.Errorf("Error writing packet to UDP: %v", err)
	}

	logrus.Debugln("Sender sent packet!")

	return nil
}
