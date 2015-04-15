package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p/udp"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Println("Begin main")
	/*
		c := make(chan id.ID, 1)
		id0 := id.New(32)

		go id0.ServeIDs(c)

		for i := 0; i < 10; i++ {
			id1 := <-c
			logrus.Println(id1.GetID())
		}
	*/

	//start the sender
	ch := make(chan int, 1)
	go func(ch chan int) {
		conn, err := udp.GetListenConn("127.0.0.1:8090")
		if err != nil {
			logrus.Errorln("Error setting up connectionless listen:", err)
		}
		defer conn.Close()

		addr, err := udp.GetAddr("127.0.0.1:8080")
		if err != nil {
			logrus.Errorln("Error getting udp address", err)
		}
		logrus.Println("udp ip is ", addr.IP)

		logrus.Debugln("Writing connectionless udp")
		some, err := conn.WriteTo([]byte("Hello there!"), addr)
		if err != nil {
			logrus.Errorln("Error writing to udp connection", err)
		}
		logrus.Println("Return value is", some)
		logrus.Debugln("Wrote to udp packet connection")
		ch <- 0 //signal exit
	}(ch)

	ch0 := make(chan int, 1)
	go func(ch chan int) {
		conn, err := udp.GetListenConn("127.0.0.1:8080")
		if err != nil {
			logrus.Errorln("Error setting up connectionless listen:", err)
		}
		defer conn.Close()

		addr, err := udp.GetAddr("127.0.0.1:8080")
		if err != nil {
			logrus.Errorln("Error getting udp address", err)
		}
		logrus.Println("udp ip is ", addr.IP)

		logrus.Debugln("Reading connectionless udp")
		some, err := conn.WriteTo([]byte("Hello there!"), addr)
		if err != nil {
			logrus.Errorln("Error writing to udp connection", err)
		}
		logrus.Println("Return value is", some)
		logrus.Debugln("Wrote to udp packet connection")
		ch <- 0 //signal exit

	}(ch0)

	//exit stuff
	<-ch
}
