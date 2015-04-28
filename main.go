package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/kodykantor/p2p-gossip/id"
	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/server"
	"github.com/kodykantor/p2p-gossip/ttl"
	"github.com/kodykantor/p2p-gossip/udp"
	"github.com/kodykantor/p2p-gossip/udp/peer"
)

//cmdGet uses REST to make a request to the peer's REST server.
//The resource requested is a GET parameter.
func cmdGet(c *cli.Context) {
	logrus.Debugln("Inside Get Command!")
	logrus.Debugln("Resource to get is ", c.String("resource"))

}

func runApp(c *cli.Context) {
	if c.Bool("verbose") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debugln("Set debug level")
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
		logrus.Errorln("Set error level")
	}

	if c.Bool("peer") { // run the peer in the foreground
		go server.ServeRest() //start the UI server

		receiveChannel := make(chan *packet.PeerPacket, 1)
		sendChannel := make(chan *packet.PeerPacket, 1)
		signalChannel := make(chan int, 1)

		mypeer := new(peer.Peer)
		err := mypeer.SetPort(8080)
		if err != nil {
			_ = fmt.Errorf("Error setting peer port: %v", err)
		}
		fmt.Println("Set peer port")

		err = mypeer.SetPacketSize(12345)
		if err != nil {
			_ = fmt.Errorf("Error setting packet size: %v", err)
		}
		logrus.Debugln("Set packet size")

		go udp.RunPeer(sendChannel, receiveChannel, signalChannel, mypeer) //start the peer

		time.Sleep(time.Second * 2)

		//BEGIN TEST
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
		sendChannel <- mypacket.(*packet.PeerPacket) //get the packet in the channel right away
		//END TEST

		//wait for SIGINT
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for sig := range c {
			fmt.Println("Received signal: ", sig)
			signalChannel <- 1
			os.Exit(1)
		}

	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gossip"
	app.Usage = "Enter a command to get resources or show resources."
	app.Action = runApp

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "verbose logging",
		},
		cli.BoolFlag{
			Name:  "p, peer",
			Usage: "run as a peer",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "get", //get a resource
			Usage:  "Request a resource from the peer network.",
			Action: cmdGet,
			Flags: []cli.Flag{ //specify the resource to download
				cli.StringFlag{
					Name:  "r, resource",
					Usage: "Specify a resource to search for or download.",
				},
			},
		},
	}

	app.Run(os.Args)

	/*
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

		<-recChan */
}
