package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var res string = "PASS"

func runTest(w http.ResponseWriter, r *http.Request) {
	i := rand.Intn(10)
	log.Printf("Endpoint Hit: runTest %d\n", i)
	if i < 5 {
		res = "FAILED"
	} else {
		res = "IN-PROCESS"
		log.Printf("sleeping for %d seconds\n", i)
		time.Sleep(time.Second * time.Duration(i))
		res = "PASS"
	}
	fmt.Fprintf(w, "ok")
}

func testResult(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, res)
	log.Println("Endpoint Hit: testResult:", res)
}

func handleRequests() {
	http.HandleFunc("/tests/run", runTest)
	http.HandleFunc("/tests/result", testResult)
	log.Fatal(http.ListenAndServe(":5003", nil))
}

func main() {
	log.Printf("service test suite is started...\n")
	handleRequests()
}
