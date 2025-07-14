package memory

import (
	"log"
	"os"
	"time"

	"carrega/daemon/models"
)

func TrackDownload(download *models.DownloadProcess, endTracker chan bool) {
	Ongoing.add(download)

	ticker := time.Tick(time.Second / 4)
	running := true
	log.Println("start")

	for running {
		select {
		case <-ticker:
			_ = updateProgress(download)
		case <-endTracker:
			log.Println("download complete")
			running = false
		}
	}

	Ongoing.remove(download)
	Finished.add(download)
}

func updateProgress(download *models.DownloadProcess) error {
	res, err := os.Stat(download.OutputDir + download.FileName)
	if err != nil {
		return err
	}

	download.Progress = formatSize(res.Size())

	log.Printf("downloading... %.1f done", download.Progress)
	return nil
}

func formatSize(entry int64) float32 {
	var shortEntry int = int(entry / 1e5)
	return float32(shortEntry) / 1e1
}
