package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var inputFilename string = "devio_articles.json"
var appPort = "8081"

func main() {

	log.Printf("Listening on :%s", appPort)
	http.HandleFunc("/", articleHandler)
	log.Fatal(http.ListenAndServe(net.JoinHostPort("", appPort), nil))
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	basePath, ok := os.LookupEnv("STORAGE_DIR")
	if !ok {
		log.Fatalf(" ENV[STORAGE_DIR] missing.")
	}
	inputFilepath := filepath.Join(basePath, inputFilename)

	content, err := readArticles(inputFilepath)
	if err != nil {
		log.Printf("\nerror reading articles from file. error: %v", err)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("Content not available"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	log.Printf("Successfully retrieved articles...")
	w.Write(content)
}

func readArticles(src string) ([]byte, error) {
	data, err := ioutil.ReadFile(src)
	return data, err
}
