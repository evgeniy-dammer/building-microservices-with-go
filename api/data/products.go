package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product model
// swagger:model
type Product struct {
	// The id for a product
	// required: true
	// min: 1
	ID int `json:"id"`

	// The name for a product
	// required: true
	Name string `json:"name" validate:"required"`

	// The description for a product
	Description string `json:"description"`

	// The price for a product
	// required: true
	Price float32 `json:"price" validate:"gt=0,required"`

	// The SKU for a product
	// required: true
	// example: SKU0001
	SKU string `json:"sku" validate:"required,sku"`

	// The date when the product was created
	CreatedOn string `json:"-"`

	// The date when the product was updated
	UpdatedOn string `json:"-"`

	// The date when the product was deleted
	DeletedOn string `json:"-"`
}

// ErrProductNotFound error product not found
var ErrProductNotFound = fmt.Errorf("Product not found")

// Products model
type Products []*Product

// ToJSON encodes an object to a JSON string
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(p)
}

// FromJSON decodes a JSON string to an object
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)

	return e.Decode(p)
}

// Validate validates a struct after deserializing JSON
func (p *Product) Validate() error {
	// create new validator
	validate := validator.New()

	// register custom SKU validation function
	validate.RegisterValidation("sku", validateSKU)

	// validate struct
	return validate.Struct(p)
}

// validateSKU is custom function for SKU validation
func validateSKU(fl validator.FieldLevel) bool {
	// string format
	reg := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")

	// searching a string with format below
	maches := reg.FindAllString(fl.Field().String(), -1)

	return len(maches) == 1
}

// GetProducts returns the list of products
func GetProducts() Products {
	return productList
}

// InsertProduct inserts product into datastore
func InsertProduct(p *Product) error {
	// create next product id
	p.ID = getNextId()

	// insert new product in datastore
	productList = append(productList, p)

	return nil
}

// UpdateProduct updates product into datastore by id
func UpdateProduct(id int, p *Product) error {
	// find product by id
	index, err := findProduct(id)

	if err != nil {
		return err
	}

	// update product in datastore
	p.ID = id
	productList[index] = p

	return nil
}

// DeleteProduct deletes product from datastore by id
func DeleteProduct(id int) error {
	// find product by id
	index, err := findProduct(id)

	if err != nil {
		return err
	}

	// delete product from datastore
	productList = append(productList[:index], productList[index+1:]...)

	return nil
}

// findProduct searchs product in datastore
func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}

	return -1, ErrProductNotFound
}

// getNextId returns next product id from datastore
func getNextId() int {
	return productList[len(productList)-1].ID + 1
}

// productList is the fixed list of products
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothly milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "cba321",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
