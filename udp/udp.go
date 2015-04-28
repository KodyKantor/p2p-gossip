//Package udp provides a peer that can send and receive
//udp packets.
package udp

import (
	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/udp/peer"
	"github.com/kodykantor/p2p-gossip/udp/receiver"
	"github.com/kodykantor/p2p-gossip/udp/sender"
)

func init() {
	logrus.Debugln("Initialized udp package.")
}

//Peer can send and receive packets.
type Peer interface {
	GetPort() int            //returns the port used by the peer
	SetPort(int) error       //sets the port used to send/receive packets
	GetPacketSize() int      //returns the max size of a udp packet
	SetPacketSize(int) error //sets the max size of a udp packet
	//RunPeer(chan *packet.Packet, chan *packet.Packet, chan int) //runs the peer until a channel signal
}

//Sender can send udp packets.
type Sender interface {
	Send(chan *packet.PeerPacket) error //sends udp packets
}

//Receiver can receive udp packets.
type Receiver interface {
	Receive(chan *packet.PeerPacket) error //receives udp packets, and places them in the channel
}

//RunPeer takes a signal channel, and channels to send and receive packets.
//Packets to be sent are read from the send channel.
//Packets read from the network are sent through the receive channel.
func RunPeer(sendChan, recChan chan *packet.PeerPacket, signal chan int, p *peer.Peer) {
	rec := receiver.New(p)
	go rec.Receive(recChan) //receive packets from the channel

	send := sender.New(p)
	go send.Send(sendChan) //send packets as they come through the channel
}
