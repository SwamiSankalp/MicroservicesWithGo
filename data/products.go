package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for an API product
// Using Struct tags we can add annotations to the fields.
type Product struct {
	ID 			int 	`json:"id"` 		// using this json tag, we can change the output field.
	Name 		string 	`json:"name"`
	Description string 	`json:"description"`
	Price 		float32 `json:"price"`
	SKU 		string 	`json:"sku"`
	CreatedOn 	string 	`json:"-"` 			// Using "-", we can omit fields from the output
	UpdatedOn 	string 	`json:"-"`
	DeletedOn 	string 	`json:"-"`
}


type Products []*Product

// A method on Product struct to convert the struct data into a valid JSON format.
// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance  than json. Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overloads of the service
func (p *Products) ToJSON(w io.Writer) error { 
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// A method on Product struct to decode the JSON format into the struct type.
func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func GetProducts() Products {
	return productList
}


func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}


func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct (id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}


func getNextId() int {
	lp := productList[len(productList) - 1]
	return lp.ID + 1
}


var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc321",
		CreatedOn: time.Now().String(),
		UpdatedOn: time.Now().String(),
	}, 
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee without milk",
		Price: 1.99,
		SKU: "abc322",
		CreatedOn: time.Now().String(),
		UpdatedOn: time.Now().String(),
	},
}