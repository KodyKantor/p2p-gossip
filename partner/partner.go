package partner

import "github.com/Sirupsen/logrus"

func init() {
	logrus.Debugln("Initialized partner package")
}

//Partner holds information about an individual gossip partner.
type Partner struct {
	Address string //IP address of the partner

}

//Partners is a list of gossip partners.
type Partners struct {
	List []*Partner // a slice of pointers to Partner structs
}
