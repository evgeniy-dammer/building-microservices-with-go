package handlers

import (
	"github.com/evgeniy-dammer/building-microservices-with-go/api/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// swagger:route PUT /products/{id} products updateProduct
// Updates product by id in database
// responses:
//
//	200: noContent

// UpdateProduct updates product by id
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")

	// extract id from path
	vars := mux.Vars(r)

	// convert string id ti int
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	// create new product object from request context
	product := r.Context().Value(KeyProduct{}).(data.Product)

	// updates an object to the datastore
	err = data.UpdateProduct(id, &product)

	// if product not founded
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	// if another error
	if err != nil {
		http.Error(rw, "Unable to update product", http.StatusInternalServerError)
		return
	}

	// if all is OK
	rw.WriteHeader(http.StatusOK)
}
