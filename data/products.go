package data

import (
	"encoding/json"
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

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList
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