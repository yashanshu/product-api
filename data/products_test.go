package data

import (
	"testing"
)

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "lolo",
		Price: 1.00,
		SKU:   "abc-asw-aa",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
