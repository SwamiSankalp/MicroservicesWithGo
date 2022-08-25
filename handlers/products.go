package handlers

import (
	"context"
	"log"
	"microservices/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a http.Handler with the given logger
func NewProducts(l*log.Logger) *Products {
	return &Products{l}
}

// ****The Below code commented code is written using the Go Standard library*****
//
// // ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface 
// func (p*Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	// handle all the request for a list of products
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		// expect the id in the uri
// 		regex := regexp.MustCompile(`/([0-9]+)`)
// 		g := regex.FindAllStringSubmatch(r.URL.Path, 1)

// 		if len(g) != 1 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		if len(g[0]) != 2 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		idString := g[0][1]
// 		id,err := strconv.Atoi(idString)
// 		if err != nil {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		p.l.Println("got id", id)
// 		p.updateProduct(id, rw, r)
// 	}

// 	// catch all
// 	// If no method is satisfied, return an error
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }

// getProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r*http.Request) {
	p.l.Println("Get Products")

	// fetch the products from the data store
	listOfProducts := data.GetProducts()
	// data,err := json.Marshal(listOfProducts) // JSON Marshal traverses all the data into a valid JSON format.
	err := listOfProducts.ToJSON(rw) // Code is bit quicker using JSON Encode method than using json.Marshal()
	if err != nil {
		http.Error(rw, "Unable to Encode to JSON", http.StatusInternalServerError)
	}
	// rw.Write(data)
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}


func (p *Products) UpdateProduct(rw  http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) 
	// the mux.Vars contains all the placeholder values passed in the URL Path
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert the Id", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT Product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct {}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}