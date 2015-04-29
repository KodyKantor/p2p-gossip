//Package server serves up REST urls for User Interfaces to the Peer.
package server

import (
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kodykantor/p2p-gossip/id"
	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/ttl"
)

func init() {
	logrus.Debugln("Initialized REST Server package.")
}

//findRequest asks peers if they have the resource provided in the GET
//parameter of the REST request.
func findRequest(c *gin.Context) {
	token := c.Request.URL.Query().Get("resource") //find the GET param
	//	c.String(http.StatusOK, token)                 //write the param back
	_ = packet.BufferizableString{Str: token}

}

//ServeRest published and handles REST endpoints.
func ServeRest(port int, sendChannel, receiveChannel chan *packet.PeerPacket) {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/request", func(c *gin.Context) {
		token := c.Request.URL.Query().Get("resource")
		bufStr := &packet.BufferizableString{Str: token}

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

		mypacket, err := pack.CreatePacket(id0, id1, ttl0, bufStr)
		if err != nil {
			logrus.Error("Error creating packet:", err)
		}
		sendChannel <- mypacket.(*packet.PeerPacket) //get the packet in the channel right away

	})

	router.Run(":" + strconv.Itoa(port))

}
