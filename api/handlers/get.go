package handlers

import (
	"github.com/evgeniy-dammer/building-microservices-with-go/api/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of products from database
// responses:
//
//	200: productsResponse

// GetProducts returns the products from data sourse
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")

	// fetch the products from datastore
	lp := data.GetProducts()

	// serialize the list of poducts to JSON
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to encode JSON", http.StatusInternalServerError)
	}
}
