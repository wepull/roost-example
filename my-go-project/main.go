package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	  
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
