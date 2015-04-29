package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/kodykantor/p2p-gossip/client"
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
	runApp(c) //check the global flags
	logrus.Debugln("Inside Get Command!")
	logrus.Debugln("Resource to get is ", c.String("resource"))

	client := client.RestClient{Hostname: "http://localhost:8080/request"}
	err := client.GetResource(c.String("resource"))
	if err != nil {
		logrus.Errorf("Error getting resource from the server: %v", err)
	}

}

func cmdPeer(c *cli.Context) {
	runApp(c)

	receiveChannel := make(chan *packet.PeerPacket, 1)
	sendChannel := make(chan *packet.PeerPacket, 1)
	signalChannel := make(chan int, 1)

	go server.ServeRest(8080, sendChannel, receiveChannel) //start the UI server
	mypeer := new(peer.Peer)
	err := mypeer.SetPort(12345)
	if err != nil {
		_ = fmt.Errorf("Error setting peer port: %v", err)
	}
	fmt.Println("Set peer port")

	err = mypeer.SetPacketSize(512)
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
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	for sig := range sigChan {
		fmt.Println("Received signal: ", sig)
		signalChannel <- 1
		os.Exit(1)
	}

}

func runApp(c *cli.Context) {
	if c.GlobalBool("verbose") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debugln("Set debug level")
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
		logrus.Errorln("Set error level")
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gossip"
	app.Usage = "Enter a command to get resources or show resources."
	app.Action = runApp //runApp is called only if no subcommands are specified

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "verbose logging",
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
		{
			Name:   "peer",
			Usage:  "Run this binary as a peer in the foreground.",
			Action: cmdPeer,
		},
	}

	app.Run(os.Args)
}
