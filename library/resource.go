package library

import "github.com/kodykantor/p2p-gossip/id"

//Resource holds information about files a peer can share.
type Resource struct {
	Name        string `json:"name"`        //name of the resource (cats.jpg)
	Description string `json:"description"` //short description of the resource
	Location    string `json:"location"`    //TODO make this a file type
	MimeType    string `json:"mimetype"`
	resourceID  id.ID  //unique id of the resource
}
