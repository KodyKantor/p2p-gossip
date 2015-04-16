package ttl

import (
	"fmt"
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

const (
	DEFAULT_TTL = 30
)

func init() {
	log.Println("Initialized ttl")
}

type TTL struct {
	ttl int //time to live
}

func New(ttl int) TTL {
	if ttl < 1 {
		ttl = DEFAULT_TTL
	}

	return TTL{ttl}
}

//SetTTL sets the ttl attribute.
func (t *TTL) SetTTL(ttl int) error {
	if ttl < 0 {
		return fmt.Errorf("Invalid time to live: %v", ttl)
	}
	t.ttl = ttl
	return nil
}

//GetTTL returns the ttl attribute.
func (t *TTL) GetTTL() int {
	return t.ttl
}
