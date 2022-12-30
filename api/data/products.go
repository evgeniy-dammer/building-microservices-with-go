package data

import (
	"context"
	"fmt"
	protos "github.com/evgeniy-dammer/building-microservices-with-go/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
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
	Price float64 `json:"price" validate:"gt=0,required"`

	// The SKU for a product
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`
}

// Products defines a slice of Product
type Products []*Product

// ProductsDB
type ProductsDB struct {
	currency protos.CurrencyClient
	log      hclog.Logger
}

// NewProductsDB
func NewProductsDB(c protos.CurrencyClient, l hclog.Logger) *ProductsDB {
	return &ProductsDB{currency: c, log: l}
}

// GetProducts returns all products from the database
func (p *ProductsDB) GetProducts(currency string) (Products, error) {
	if currency == "" {
		return productList, nil
	}

	rate, err := p.getRate(currency)

	if err != nil {
		p.log.Error("[ERROR] fetching exchange rate", "currency", currency, "error", err)

		return nil, err
	}

	pr := Products{}

	for _, p := range productList {
		np := *p
		np.Price = np.Price * rate

		pr = append(pr, &np)
	}

	return pr, nil
}

// GetProductById returns a single product which matches the id from the database.
func (p *ProductsDB) GetProductById(id int, currency string) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	if currency == "" {
		return productList[i], nil
	}

	rate, err := p.getRate(currency)

	if err != nil {
		p.log.Error("[ERROR] fetching exchange rate", "currency", currency, "error", err)

		return nil, err
	}

	np := *productList[i]
	np.Price = np.Price * rate

	return &np, nil
}

func (p *ProductsDB) getRate(destination string) (float64, error) {
	// get the exchange rate
	rr := protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[destination]),
	}

	resp, err := p.currency.GetRate(context.Background(), &rr)

	return resp.Rate, err
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func (p *ProductsDB) UpdateProduct(prod Product) error {
	i := findIndexByProductID(prod.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	// update the product in the DB
	productList[i] = &prod

	return nil
}

// AddProduct adds a new product to the database
func (p *ProductsDB) AddProduct(prod Product) {
	// get the next id in sequence
	maxID := productList[len(productList)-1].ID
	prod.ID = maxID + 1
	productList = append(productList, &prod)
}

// DeleteProduct deletes a product from the database
func (p *ProductsDB) DeleteProduct(id int) error {
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
