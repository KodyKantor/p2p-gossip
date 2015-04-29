package server

import (
	"testing"
	"time"

	"github.com/kodykantor/p2p-gossip/client"
	"github.com/kodykantor/p2p-gossip/packet"
)

func TestServerPing(t *testing.T) {
	t.Log("Testing Ping functionality")
	go ServeRest(8080, nil, nil) //run in the background

	time.Sleep(time.Millisecond * 500) //sleep for a half second
	cli := client.RestClient{Hostname: "http://localhost:8080"}

	str, err := cli.Ping()
	if err != nil {
		t.Errorf("Error pinging server: %v", err)
	}

	t.Logf("Server said: %v", str)

}

func TestServerRequest(t *testing.T) {
	t.Log("Testing 'Find' functionality")
	mych := make(chan *packet.PeerPacket, 1)
	go ServeRest(8070, mych, mych)

	time.Sleep(time.Millisecond * 500) //sleep .5 seconds
	cli := client.RestClient{Hostname: "http://localhost:8070"}

	str, err := cli.Resource("cats.jpg")
	if err != nil {
		t.Errorf("Error requesting resource from server: %v", err)
	}

	t.Logf("Server said: %v", str)
}
