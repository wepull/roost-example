package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var outputFilename string = "devio_articles_"

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

func getFilePath(tag string) string {
	basePath, ok := os.LookupEnv("STORAGE_DIR")
	if !ok {
		log.Fatalf(" ENV[STORAGE_DIR] missing.")
	}

	outputPath := filepath.Join(basePath, outputFilename+tag+".json")
	log.Println("%s", outputPath)
	return outputPath
}

// save articles to local filesystem
func save(data []byte, tag string) {

	outputPath := getFilePath(tag)
	log.Println("%s", outputPath)
	err := ioutil.WriteFile(outputPath, data, 0644)
	if err != nil {
		log.Printf("error saving output in file at: %s, error: %v", outputPath, err)
	} else {
		log.Printf("successfully saved articles to file. %s", outputPath)
	}
}

// articleHandler handles request to fetch articles
func fetchData(tags []string) {

	for _, tag := range tags {
		reqQueryParam := ""

		if tag != "" {
			reqQueryParam = fmt.Sprintf("?tag=%s", tag)
		}
		log.Printf("Request with tag: %+v\n", tag)

		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://dev.to/api/articles%s", reqQueryParam), nil)

		resp, err := reqDevTO(req)
		if err != nil {
			log.Printf("unable to get articles. Error: %v", err)
			return
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("error reading response from dev.to api. Error: %v", err)
		}

		var d bytes.Buffer
		json.Indent(&d, data, "", "\t")

		save(d.Bytes(), tag)
	}

}

// articleHandler handles request to fetch articles
func articleHandler(w http.ResponseWriter, r *http.Request) {
	reqQueryParam := ""
	tag := r.URL.Query().Get("tag")
	if tag != "" {
		reqQueryParam = fmt.Sprintf("?tag=%s", tag)
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
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response from dev.to api. Error: %v", err)
	}
	articles := data

	// Indent only json data
	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		var d bytes.Buffer
		json.Indent(&d, data, "", "\t")
		articles = d.Bytes()
	}

	w.Write(articles)

	save(articles, tag)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/articles", 301)
}

func cleanHandler(w http.ResponseWriter, r *http.Request) {
	outputPath := getFilePath(r.URL.Query().Get("tag"))

	if err := os.Remove(outputPath); err != nil {
		log.Printf("Unable to delete article store file. error: %v", err)
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Unable to delete article store file. error: %v", err)))
		return
	}
	log.Println("Sucessfully deleted article store")
	w.Write([]byte("Sucessfully deleted article store"))
}

func serveHandler(w http.ResponseWriter, r *http.Request) {
	outputPath := getFilePath(r.URL.Query().Get("tag"))

	content, err := readArticles(outputPath)
	if err != nil {
		log.Printf("\nerror reading articles from file. Request received from address: %s, error: %v", r.RemoteAddr, err)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("Content not available"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	log.Printf("Successfully retrieved articles... Request received from address: %s", r.RemoteAddr)
	w.Write(content)
}

func readArticles(src string) ([]byte, error) {
	data, err := ioutil.ReadFile(src)
	return data, err
}

func main() {
	log.Printf("Listening on :8080")

	// http.HandleFunc("/", rootHandler)
	http.HandleFunc("/articles", articleHandler)
	http.HandleFunc("/clean", cleanHandler)
	http.HandleFunc("/serve", serveHandler)

	log.Printf("Going to call fetchData")
	fetchData([]string{"kubernetes", "docker", "golang", "roost", "vertx"})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
