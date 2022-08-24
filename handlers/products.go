package handlers

import (
	"log"
	"microservices/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products {
	return &Products{l}
}

func (p*Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}


func (p *Products) getProducts(rw http.ResponseWriter, r*http.Request) {
	p.l.Println("Get Products")
	listOfProducts := data.GetProducts()
	// data,err := json.Marshal(listOfProducts) // JSON Marshal traverses all the data into a valid JSON format.
	err := listOfProducts.ToJSON(rw) // Code is bit quicker using JSON Encode method than using json.Marshal()
	if err != nil {
		http.Error(rw, "Unable to Encode to JSON", http.StatusInternalServerError)
	}
	// rw.Write(data)
}