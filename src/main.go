package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func localWatchHandler(watcher *fsnotify.Watcher, file string) {

	minio := Minio{"localhost:9000", os.Getenv("MINIO_ACCESS_ID"), os.Getenv("MINIO_ACCESS_KEY"), false, "kp-sync", context.Background()}
	client := minio.initMinio()

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Println("event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name)
				fmt.Println("writing to server")
				minio.uploadfile(client, "file", file, "application/octet-stream")
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
func filewatch(file string) *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	//	done := make(chan bool)

	err = watcher.Add(file)
	if err != nil {
		log.Fatal(err)
	}
	//<-done
	return watcher
}
func main() {
	filePath := "/tmp/foo"
	fmt.Println("Startup!")

	// Start file watch
	watcher := filewatch(filePath)
	go localWatchHandler(watcher, filePath)

	select {}

}
