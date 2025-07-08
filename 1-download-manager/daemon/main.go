package main

import (
	"log"

	"carrega/daemon/listener/unix"
	"carrega/daemon/listener/unixgram"
)

func main() {
	log.Println("running")
	go unixgram.StartListener()
	go unix.StartServer()
	for {
	}
}
