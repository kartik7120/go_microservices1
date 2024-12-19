package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Price     float32 `json:"price"`
	SKU       string  `json:"sku"` // what is SKU? Stock Keeping Unit - a unique identifier for each distinct product and service that can be purchased.
	CreatedOn string  `json:"-"`   // "-" means that this field will not be marshalled or unmarshalled from JSON. what does marshalling and unmarshalling mean? Marshalling is the process of transforming the memory representation of an object to a data format suitable for storage or transmission. Unmarshalling is the reverse process.
	UpdatedOn string  `json:"-"`   // "-" means that this field will not be marshalled or unmarshalled from JSON
	DeletedOn string  `json:"-"`   // "-" means that this field will not be marshalled or unmarshalled from JSON
}

type Products []*Product

// Products is a collection of Product

var productList = []*Product{
	{
		ID:        1,
		Name:      "Latte",
		Desc:      "Frothy milky coffee",
		Price:     2.45,
		SKU:       "abc323",
		CreatedOn: "2021-01-01T10:00:00Z",
		UpdatedOn: "2021-01-01T10:00:00Z",
	},
	{
		ID:        2,
		Name:      "Espresso",
		Desc:      "Short and strong coffee without milk",
		Price:     1.99,
		SKU:       "fjd34",
		CreatedOn: "2021-01-01T10:00:00Z",
		UpdatedOn: "2021-01-01T10:00:00Z",
	},
}

// In order to access the productList from outside the package, we can abstract the access to the list by creating a function that returns the list. This is a common pattern in Go to provide access to private variables.

func GetProducts() Products {
	return productList
}

// Using an encorder to encode the products into a json output
// Why use an encoder instead of json.marshal ? The encoder is more efficient and can be used to encode the data directly to an io.Writer. This is useful when you want to write the data directly to a network connection or a file without having to store the entire JSON in memory.

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
