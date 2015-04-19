package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/id"
	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/ttl"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Println("Begin main")

	var myid id.ID
	var myttl ttl.TTL
	var pack packet.Packet

	myid = new(id.PeerID)
	myttl = new(ttl.PeerTTL)
	pack = new(packet.PeerPacket)

	myid.SetLength(32)
	ch := make(chan id.ID, 1)

	go myid.ServeIDs(ch)

	id0 := <-ch
	id1 := <-ch
	ttl0, err := myttl.CreateTTL(32)
	if err != nil {
		logrus.Error("Error getting ttl:", err)
	}

	mypacket, err := pack.CreatePacket(id0, id1, ttl0)
	if err != nil {
		logrus.Error("Error creating packet:", err)
	}

	logrus.Println("Mypacket is:", mypacket.GetBuffer())

}
