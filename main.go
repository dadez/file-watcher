package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
)

const directoryWatch = "/tmp/testdir"

func postFile(f string) {

	log.Println("upload called for file:", f)

	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	url := "https://httpbin.org/post"
	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(file))
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					postFile(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(directoryWatch)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("watching directory: ", directoryWatch)
	<-done
}
