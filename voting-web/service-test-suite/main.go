package main

import (
	"fmt"
	"log"
	"net/http"
)

func runTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
	fmt.Println("Endpoint Hit: runTest")
}

func testResult(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PASS")
	fmt.Println("Endpoint Hit: testResult")
}

func handleRequests() {
	http.HandleFunc("/tests/run", runTest)
	http.HandleFunc("/tests/result", testResult)
	log.Fatal(http.ListenAndServe(":5003", nil))
}

func main() {
	log.Printf("service test suite is started...")
	handleRequests()
}
