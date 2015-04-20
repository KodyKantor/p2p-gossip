//Package udp provides a peer that can send and receive
//udp packets.
package udp

import (
	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/packet"
)

var log = logrus.New()

func init() {
	log.Println("Initialized udp")
}

//Peers can send and receive packets.
type Peer interface {
	GetPort() int                         //returns the port used by the peer
	SetPort(int) error                    //sets the port used to send/receive packets
	GetPacketSize() int                   //returns the max size of a udp packet
	SetPacketSize(int) error              //sets the max size of a udp packet
	RunPeer(chan int, chan int, chan int) //runs the peer until a channel signal
}

//Senders can send udp packets.
type Sender interface {
	Send(chan packet.Packet) error //sends udp packets
}

//Receivers can receive udp packets.
type Receiver interface {
	Receive(chan packet.Packet) error //receives udp packets, and places them in the channel
}
