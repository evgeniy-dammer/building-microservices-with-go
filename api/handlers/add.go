package handlers

import (
	"github.com/evgeniy-dammer/building-microservices-with-go/api/data"
	"net/http"
)

// swagger:route POST /products products addProduct
// Inserts a new product in database
// responses:
//
//	201: productsRequest

// AddProduct insert new product to the datastore
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	// create new product object from request context
	product := r.Context().Value(KeyProduct{}).(data.Product)

	// insert an object to the datastore
	err := data.InsertProduct(&product)

	if err != nil {
		http.Error(rw, "Unable to add product", http.StatusInternalServerError)
		return
	}

	// if all is OK
	rw.WriteHeader(http.StatusCreated)
}
