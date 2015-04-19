package udp

import (
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Println("Initialized udp")
}
