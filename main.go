package main

import (
	"fmt"
	"time"

	"github.com/Sanchit011/blockchaingo/network"
)

func main() {
	fmt.Println("hello")

	trLocal := network.NewLocalTransport("Local")
	trRemote := network.NewLocalTransport("Remote")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func () {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("message from remote"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}