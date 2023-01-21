package data

import "testing"


func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name: "Victor",
		Price: 2.00,
		SKU: "abs-abs-sjd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}