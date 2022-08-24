package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewHello (l *log.Logger) *Hello {
	return &Hello{l}
}

	// http.ResponseWriter is a interface to construct an HTTP response 
	// Ex: It can write to headers, statuscodes, response body and so on
	// http.Request represents an HTTP request
	// It contains info like PATH, METHODS, BODY, HTTP version and so on
// ServeHTTP implements the go http.Handler interface	
func (h*Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// http.HandleFunc registers a function to a path on the defaultServeMux 
		// => http.DefaultServeMux.HandleFunc("/", func)
		// DefaultServeMux is a http handler and every thing related to the go server is a http handler
		h.l.Println("Hello World")

		// read the body
		d,err := ioutil.ReadAll(r.Body) // ioutil reads all the data passed through the request
		if err != nil { // Check if error and write an error message to the request
			// rw.WriteHeader(http.StatusBadRequest) // WriteHeader allows to specify HTTP StatusCode
			// rw.Write([]byte("Oops!"))
			http.Error(rw, "Ooops!", http.StatusBadRequest) 
			// http has a standard Error interface to handle everything related to the errors
			return
		}
		
		// write the response
		fmt.Fprintf(rw, "Hello %s\n", d) // fmt.Fprintf allows to write responses
}

// “*” says, “you are declaring that this variable holds a memory address to a string, or int or whatever type follows “*”. 
// For example, “var a *int” declares that the variable “a” holds a memory address(pointer) to an int datatype.

