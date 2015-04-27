package server

import (
	"testing"

	"github.com/kodykantor/p2p-gossip/client"
)

func TestServeRest(t *testing.T) {
	go ServeRest() //run in the background

	cli := client.RestClient{Hostname: "http://localhost:8080"}

	str, err := cli.Ping()
	if err != nil {
		t.Errorf("Error pinging server: %v", err)
	}

	t.Logf("Server said: %v", str)

}
