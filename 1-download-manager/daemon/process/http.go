package process

import (
	"io"
	"log"
	"net/http"
	"os"

	"carrega/daemon/memory"
	"carrega/daemon/models"
)

func Download(ops *models.DownloadProcess) error {
	log.Println("downloading", ops.Url)
	res, err := http.Get(ops.Url)
	if err != nil {
		return err
	}

	os.MkdirAll(ops.OutputDir, 0755)
	f, err := os.Create(ops.OutputDir + ops.FileName)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	endTracker := make(chan bool)
	go memory.TrackDownload(ops, endTracker)

	if _, err := io.Copy(f, res.Body); err != nil {
		return err
	}

	endTracker <- true
	return nil
}
