package main

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/id"
	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/ttl"
	"github.com/kodykantor/p2p-gossip/udp/peer"
	"github.com/kodykantor/p2p-gossip/udp/receiver"
	"github.com/kodykantor/p2p-gossip/udp/sender"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Println("Begin main")

	//Create a packet
	var myid id.ID
	var myttl ttl.TTL
	var pack packet.Packet

	myid = id.NewID()
	myttl = ttl.NewTTL()
	pack = packet.NewPacket()

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

	logrus.Debugln("Mypacket is:", mypacket.GetBufferization())

	mypeer := new(peer.Peer)
	err = mypeer.SetPort(8080)
	if err != nil {
		fmt.Errorf("Error setting peer port: %v", err)
	}
	fmt.Println("Set peer port")

	err = mypeer.SetPacketSize(12345)
	if err != nil {
		fmt.Errorf("Error setting packet size: %v", err)
	}
	logrus.Debugln("Set packet size")

	recChan := make(chan *packet.PeerPacket, 1)
	go func(chan *packet.PeerPacket) {

		//Create a 'receiver'
		logrus.Debugln("Starting to receive packets...")
		myreceiver := receiver.New(mypeer)
		logrus.Debugln("Created receiver")
		err = myreceiver.Receive(recChan)
		if err != nil {
			fmt.Errorf("Error receiving packet: %v", err)
		}

		recdPacket := <-recChan
		logrus.Debugln("Received packet!")
		logrus.Debugf("recdPacket: id0: %v, id1: %v, ttl: %v\n", recdPacket.Id0.GetBytes(), recdPacket.Id1.GetBytes(), recdPacket.Ttl.GetBytes())
		recChan <- packet.NewPacket()
	}(recChan)

	//Create a 'sender'
	sendChan := make(chan *packet.PeerPacket, 1)
	sendChan <- mypacket.(*packet.PeerPacket) //get the packet in the channel right away

	time.Sleep(time.Second * 2)
	sender := sender.New(mypeer)
	err = sender.Send(sendChan) // send the single packet (should be read right away)
	if err != nil {
		fmt.Errorf("Error sending packet: %v", err)
	}

	logrus.Debugln("Sent packet!")

	<-recChan
}
