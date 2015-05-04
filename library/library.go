//Package library is a package that keeps track of resources
//available to a user. A library consists of resources that a peer can offer,
//and resources that a peer has received.
package library

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/id"
)

func init() {
	logrus.Debugln("Initialized library package")
}

//CreateResourceFromJSON take a string representation of json data, and creates
// a Resource from it.
func CreateResourceFromJSON(buf []byte) (*Resource, error) {
	if buf == nil {
		return &Resource{}, fmt.Errorf("Json buffer is nil")
	}
	if len(buf) == 0 {
		return &Resource{}, fmt.Errorf("Json buffer is zero-length")
	}

	dest := new(Resource)

	err := json.Unmarshal(buf, dest)
	if err != nil {
		return &Resource{}, fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	return dest, nil
}

//LoadResourcesFrom takes a path to a file that contains JSON descriptions
//of resources. This function will create a slice of pointers to resources.
func LoadResourcesFrom(path string) ([]*Resource, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}
	logrus.Debugln("File said:", string(data))

	var dest []*Resource
	err = json.Unmarshal(data, &dest) //convert bytes to a slice of structs
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling json: %v", err)
	}
	logrus.Debugln("JSON is:", string(data))

	return dest, nil
}

//AssignRandomIDs takes a slice of resources, and assigns each a random ID.
func AssignRandomIDs(slice []*Resource) error {
	if slice == nil {
		return fmt.Errorf("Provided slice is nil")
	}
	if len(slice) == 0 {
		return fmt.Errorf("Provided slice is zero-length")
	}
	myid := id.NewID()
	ch := make(chan id.ID, 1)

	go myid.ServeIDs(ch) //start serving IDs in the background
	for _, val := range slice {
		val.resourceID = <-ch //read from the channel, assign to resource
		logrus.Debugln("Applied random ID to resource: ", val)
	}

	return nil
}
