package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var outputFilename string = "devio_articles.json"

func reqDevTO(r *http.Request) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 30,
	}

	// apiKey, ok := os.LookupEnv("DEVIO_API_KEY")
	// if !ok {
	// 	log.Fatal("missing API Key ENV[DEVIO_API_KEY]")
	// }

	// r.Header.Set("api-key", apiKey)

	resp, err := client.Do(r)
	if err != nil {
		log.Printf("error requesting dev.to. error: %v", err)
		return nil, err
	}
	return resp, err
}

func save(data []byte) {
	basePath, ok := os.LookupEnv("STORAGE_DIR")
	if !ok {
		log.Fatalf(" ENV[STORAGE_DIR] missing.")
	}

	outputPath := filepath.Join(basePath, outputFilename)
	err := ioutil.WriteFile(outputPath, data, 0644)
	if err != nil {
		log.Printf("error saving output in file at: %s, error: %v", outputPath, err)
	} else {
		log.Printf("successfully saved articles to file. %s", outputPath)
	}
}

// articleHandler handles request to fetch articles
func articleHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	queryParams := r.URL.Query()
	reqQueryParam := ""
	tag := queryParams["tag"]
	if len(tag) > 0 {
		reqQueryParam = fmt.Sprintf("?tag=%s", tag[0])
	}
	log.Printf("Request with tag: %+v\n", tag)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://dev.to/api/articles%s", reqQueryParam), nil)

	resp, err := reqDevTO(req)
	if err != nil {
		log.Printf("unable to get articles. Error: %v", err)
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("No articles retrieved. Reason: %v", err.Error())))
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response from dev.to api. Error: %v", err)
	}
	var d bytes.Buffer
	json.Indent(&d, data, "", "\t")
	// log.Printf("%v", string(d.Bytes()))

	w.Write(d.Bytes())
	save(d.Bytes())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/articles", 301)
}

func main() {
	flag.Parse()
	log.Printf("Listening on :8080")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/articles", articleHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
