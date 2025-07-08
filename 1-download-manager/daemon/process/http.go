package process

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"carrega/daemon/models"
)

func Download(ops *models.DownloadOptions) error {
	log.Println("downloading", ops.Url)
	res, err := http.Get(ops.Url)
	if err != nil {
		return err
	}

	f, err := os.Create(ops.OutputDir + ops.FileName)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	endTracker := make(chan bool)

	go func() {
		ticker := time.Tick(time.Millisecond)
		for {
			select {
			case <-ticker:
				log.Println("downloading...")
			case <-endTracker:
				log.Println("download complete")
				return
			}
		}
	}()

	if _, err := io.Copy(f, res.Body); err != nil {
		return err
	}

	endTracker <- true

	/* give the tracker time to send a final message */
	time.Sleep(time.Nanosecond)

	return nil
}
