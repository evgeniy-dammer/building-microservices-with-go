package data

import (
	"fmt"
)

// ErrProductNotFound error product not found
var ErrProductNotFound = fmt.Errorf("product not found")

// Product model
// swagger:model
type Product struct {
	// The id for a product
	// required: true
	// min: 1
	ID int `json:"id"`

	// The name for a product
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// The description for a product
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// The price for a product
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0,required"`

	// The SKU for a product
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`
}

// Products defines a slice of Product
type Products []*Product

// GetProducts returns all products from the database
func GetProducts() Products {
	return productList
}

// GetProductById returns a single product which matches the id from the database.
func GetProductById(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func UpdateProduct(p Product) error {
	i := findIndexByProductID(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	// update the product in the DB
	productList[i] = &p

	return nil
}

// AddProduct adds a new product to the database
func AddProduct(p Product) {
	// get the next id in sequence
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}
