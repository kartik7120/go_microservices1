package data

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

// Products is a collection of Product

var productList = []*Product{
	&Product{
		ID:        1,
		Name:      "Latte",
		Desc:      "Frothy milky coffee",
		Price:     2.45,
		SKU:       "abc323",
		CreatedOn: "2021-01-01T10:00:00Z",
		UpdatedOn: "2021-01-01T10:00:00Z",
	},
	&Product{
		ID:        2,
		Name:      "Espresso",
		Desc:      "Short and strong coffee without milk",
		Price:     1.99,
		SKU:       "fjd34",
		CreatedOn: "2021-01-01T10:00:00Z",
		UpdatedOn: "2021-01-01T10:00:00Z",
	},
}
