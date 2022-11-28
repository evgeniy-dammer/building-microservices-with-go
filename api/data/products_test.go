package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "sjdcjsdn",
		Price: 1.0,
		SKU:   "abc-def-ghi",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
