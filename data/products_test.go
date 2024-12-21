package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Latte",
		Price: 1.00,
		// SKU:   "abs",
		SKU: "abs-abc-def",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
