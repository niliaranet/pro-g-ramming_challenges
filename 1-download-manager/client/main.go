package main

import (
	"log"
	"time"

	"carrega/client/signal"
)

func main() {
	/*
		for range 4 {
			err := signal.SendDownload("download http://niliara.net/mothracompat.png")
			if err != nil {
				log.Println(err)
				return
			}
			time.Sleep(time.Second)
		}

		out, err := signal.Send("download http://niliara.net/mothracompat.png")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(out)
	*/

	err := signal.SendDownload("download https://discord.com/api/download?platform=linux")
	if err != nil {
		log.Println(err)
		return
	}
	time.Sleep(time.Second)
}
