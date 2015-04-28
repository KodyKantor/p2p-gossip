package server

import (
	"testing"
	"time"

	"github.com/kodykantor/p2p-gossip/client"
)

func TestServerPing(t *testing.T) {
	go ServeRest() //run in the background

	time.Sleep(time.Second * 1) //sleep for a second
	cli := client.RestClient{Hostname: "http://localhost:8080"}

	str, err := cli.Ping()
	if err != nil {
		t.Errorf("Error pinging server: %v", err)
	}

	t.Logf("Server said: %v", str)

}

func TestServerRequest(t *testing.T) {
	go ServeRest()

	time.Sleep(time.Second * 1)
	cli := client.RestClient{Hostname: "http://localhost:8080"}

	str, err := cli.Resource("cats.jpg")
	if err != nil {
		t.Errorf("Error requesting resource from server: %v", err)
	}

	t.Logf("Server said: %v", str)
}
