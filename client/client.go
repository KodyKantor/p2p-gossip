package client

import (
	"fmt"

	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/jmcvetta/napping"
)

func init() {
	logrus.Debugln("Initialized REST Client package.")
}

//RestClient makes REST requests as part of the CLI UI.
type RestClient struct {
	Hostname string //hostname of REST server
}

//GetResource tells the peer to request the resource in the 'resource' parameter.
func (r *RestClient) GetResource(resource string) error {
	_, err := napping.Get(r.Hostname, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("Error making request for resource: %v", err)
	}
	return nil
}

//Ping tests basic connection to the server.
//Returns what the server said (as text).
func (r *RestClient) Ping() (string, error) {
	resp, err := napping.Get(r.Hostname+"/ping", nil, nil, nil)
	if err != nil {
		return "", fmt.Errorf("Error reaching /ping: %v", err)
	}
	if resp.Status() != http.StatusOK {
		return "", fmt.Errorf("Did not get HTTP OK signal from server. Instead we got: %v", resp.Status())
	}

	return resp.RawText(), nil
}

//Resource sends a request to the server for the provided resource string.
//Returns what the server said (as text).
func (r *RestClient) Resource(resource string) (string, error) {
	argMap := napping.Params{}
	argMap["resource"] = resource
	resp, err := napping.Get(r.Hostname+"/request", &argMap, nil, nil)
	if err != nil {
		return "", fmt.Errorf("Error reaching /request: %v", err)
	}
	if resp.Status() != http.StatusOK {
		return "", fmt.Errorf("Did not get HTTP OK signal from server. Instead we got: %v", resp.Status())
	}

	return resp.RawText(), nil
}
