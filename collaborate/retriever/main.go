package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var appPort = "8081"

func main() {

	log.Printf("Listening on :%s", appPort)
	http.HandleFunc("/", articleHandler)
	log.Fatal(http.ListenAndServe(net.JoinHostPort("", appPort), nil))
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	fetcherEndpoint := os.Getenv("FETCH_SERVICE")
	if fetcherEndpoint == "" {
		log.Fatal("Missing fetcher service endpoint. Set ENV[FETCH_SERVICE] to continue")
	}
	
	endpoint := "http://" + fetcherEndpoint + "/serve?tag=" + r.URL.Query().Get("tag")
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	resp, err := doRequest(req)
	if err != nil {
		log.Printf("unable to get articles from the service. Error: %v", err)
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("No articles retrieved. Reason: %v", err.Error())))
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response from service. Error: %v", err)
	}

	w.Write(data)
}

func doRequest(r *http.Request) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(r)
	if err != nil {
		log.Printf("error requesting to fetcher service at endpoint. error: %v", err)
		return nil, err
	}
	return resp, err
}
