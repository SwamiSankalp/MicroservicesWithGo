package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye (l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g*GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) { 
	// Not using "*" on ResponseWriter because its an interface
	g.l.Println("Good Bye")
	d,err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops!", http.StatusBadRequest) 
			return
		}
		fmt.Fprintf(rw, "Good Bye %s\n", d)
}