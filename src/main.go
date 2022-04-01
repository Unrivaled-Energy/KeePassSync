package main

import (
	"log"
	"fmt"

	"github.com/fsnotify/fsnotify"
)
func localWatchHandler(watcher *fsnotify.Watcher){
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					fmt.Println("write to server")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
}
func filewatch(file string)(*fsnotify.Watcher){
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
  fmt.Println("Startup!")
// Start file watch
watcher := filewatch("/tmp/foo")
go localWatchHandler(watcher)


select{}

}