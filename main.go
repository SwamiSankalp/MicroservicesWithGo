package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// http.ResponseWriter is a interface to construct an HTTP response 
	// Ex: It can write to headers, statuscodes, response body and so on
	// http.Request represents an HTTP request
	// It contains info like PATH, METHODS, BODY, HTTP version and so on
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { 
		// http.HandleFunc registers a function to a path on the defaultServeMux 
		// => http.DefaultServeMux.HandleFunc("/", func)
		// DefaultServeMux is a http handler and every thing related to the go server is a http handler
		log.Println("Hello World")
		d,err := ioutil.ReadAll(r.Body) // ioutil reads all the data passed through the request
		if err != nil { // Check if error and write an error message to the request
			// rw.WriteHeader(http.StatusBadRequest) // WriteHeader allows to specify HTTP StatusCode
			// rw.Write([]byte("Oops!"))
			http.Error(rw, "Ooops!", http.StatusBadRequest) 
			// http has a standard Error interface to handle everything related to the errors
			return
		}
		fmt.Fprintf(rw, "Data %s\n", d) // fmt.Fprintf allows to write responses
	})


	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Good Bye!")
	})

	http.ListenAndServe(":8080", nil) // Binding the web server to the PORT
}