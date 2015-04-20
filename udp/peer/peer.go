//Peer implements the udp Peer interface.
package udp

import (
	"fmt"
	"github.com/kodykantor/p2p-gossip/packet"
)

//TODO change hard-coded checks to constants.

//Peer implements the Peer interface
type Peer struct {
	port       int      //port for communication
	packetsize int      //max size of packets
	signal     chan int //channel through which to send signals
}

func init() {
	fmt.Println("Initialized peer")
}

//GetPort returns the peer's port.
func (p *Peer) GetPort() int {
	return p.port
}

//SetPort sets the send/receive port to provided port number.
func (p *Peer) SetPort(port int) error {
	if port < 80 || port > 65565 {
		return fmt.Errorf("Invalid port: %v", port)
	}
	p.port = port
	return nil
}

//GetPacketSize returns the packetsize attribute.
func (p *Peer) GetPacketSize() int {
	return p.packetsize
}

//SetPacketSize sets the maximum packet size.
func (p *Peer) SetPacketSize(packetsize int) error {
	if packetsize < 1 || packetsize > 1000 {
		return fmt.Errorf("Invalid packet size: %v", packetsize)
	}

	p.packetsize = packetsize
	return nil
}

//RunPeer takes a signal channel, and channels to send and receive packets.
//Packets to be sent are read from the send channel.
//Packets read from the network are sent through the receive channel.
func (p *Peer) RunPeer(signal chan int, send, receive chan packet.Packet) {
}
