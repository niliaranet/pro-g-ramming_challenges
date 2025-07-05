package process

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"carrega/daemon/models"
)

func Download(ops *models.DownloadOptions) error {
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
				fmt.Println("hello")
			case <-endTracker:
				fmt.Println("end")
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
