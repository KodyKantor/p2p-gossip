//Package peer implements the udp Peer interface.
package peer

import "fmt"

const (
	minimumPortValue  = 80
	maximumPortValue  = 65535
	maximumPacketSize = 1000
	minimumPacketSize = 1
)

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
	if port < minimumPortValue || port > maximumPortValue {
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
	if packetsize < minimumPacketSize || packetsize > maximumPacketSize {
		return fmt.Errorf("Invalid packet size: %v", packetsize)
	}

	p.packetsize = packetsize
	return nil
}
