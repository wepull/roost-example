package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var inputFilename string = "devio_articles.json"
var refreshInterval = 5 * time.Second
var shortRefreshInterval = 3 * time.Second

func main() {
	basePath, ok := os.LookupEnv("STORAGE_DIR")
	if !ok {
		log.Fatalf(" ENV[STORAGE_DIR] missing.")
	}
	inputFilepath := filepath.Join(basePath, inputFilename)

	log.Printf("I would refresh in every %s to read content from file %s", refreshInterval.String(), inputFilepath)
	for {
		data, err := ioutil.ReadFile(inputFilepath)
		if err != nil {
			log.Printf("\nerror reading articles from file. error: %v", err)
			log.Printf("\nsleeping for %s", shortRefreshInterval.String())
			time.Sleep(shortRefreshInterval)
		}
		log.Printf("%v", string(data))
		time.Sleep(refreshInterval)
	}
}
