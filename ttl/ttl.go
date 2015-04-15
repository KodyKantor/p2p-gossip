package ttl

import (
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Println("Initialized ttl")
}

type TTL struct {
	ttl int //time to live
}
