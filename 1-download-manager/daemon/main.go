package main

import (
	"fmt"

	"carrega/daemon/listener"
)

func main() {
	fmt.Println("running")
	listener.StartListener()
}
