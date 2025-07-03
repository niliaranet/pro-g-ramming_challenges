package main

import (
	"log"

	"carrega/models"
	"carrega/process"
)

func main() {
	// url := "https://mirrors.ircam.fr/pub/zorinos-isos/17/Zorin-OS-17.3-Core-64-bit-r2.iso"
	url := "http://niliara.net/mothracompat.png"

	var ops models.DownloadOptions
	err := process.Download(ops.From(url))
	if err != nil {
		log.Fatal(err)
	}
}
