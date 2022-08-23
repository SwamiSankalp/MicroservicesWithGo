package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		// http.HandleFunc registers a function to a path on the defaultServeMux => // http.DefaultServeMux.HandleFunc("/", func)
		// DefaultServeMux is a http handler and every thing related to the go server is a http handler
		log.Println("Hello World")
	})
	http.ListenAndServe(":8080", nil) // Binding the web server to the PORT
}