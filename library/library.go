//Package library is a package that keeps track of resources
//available to a user. A library consists of resources that a peer can offer,
//and resources that a peer has received.
package library

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
)

func init() {
	logrus.Debugln("Initialized library package")
}

//CreateResourceFromJSON take a string representation of json data, and creates
// a Resource from it.
func CreateResourceFromJSON(jsn string) (*Resource, error) {
	if len(jsn) == 0 {
		return &Resource{}, fmt.Errorf("Json string is zero-length")
	}
	dest := new(Resource)
	reader := strings.NewReader(jsn)
	buf := make([]byte, len(jsn))

	_, err := reader.Read(buf)
	if err != nil {
		return &Resource{}, fmt.Errorf("Error reading string into byte slice: %v", err)
	}

	err = json.Unmarshal(buf, dest)
	if err != nil {
		return &Resource{}, fmt.Errorf("Error unmarshalling JSON: %v", err)
	}
	return dest, nil
}

//LoadResourcesFrom takes a path to a file that contains JSON descriptions
//of resources. This function will create a slice of pointers to resources.
//func LoadResourcesFrom(path string) ([]*Resource, error) {

//}
