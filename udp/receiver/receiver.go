//Receiver implements the udp Receiver interface.
package udp

import (
	"fmt"
)

func init() {
	fmt.Println("Initialized receiver")
}

type Receiver struct {
	peer Peer //nested struct
}
