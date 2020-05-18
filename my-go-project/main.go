package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"path/filepath"
)

var appPath = "/app"
var appPort = "8080"

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	absPath, _ := filepath.Abs(appPath)
	html, err := template.ParseFiles(filepath.Join(absPath, "view.html"))
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func main() {
	log.Printf("Serving at port: %s.", appPort)
	log.Printf("Deployment port may differ from application port. If running from RDE, access app in browser at: %s", "http://roost-master:30047")
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(net.JoinHostPort("", appPort), nil)
}
