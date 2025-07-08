package main

import (
	"carrega/client/signal"
	"fmt"
)

func main() {
	out, err := signal.Send("download http://niliara.net/mothracompat.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out)
}
