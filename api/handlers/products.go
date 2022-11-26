package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/evgeniy-dammer/building-microservices-with-go/api/data"
	"github.com/gorilla/mux"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// KeyProduct
type KeyProduct struct {
}

// NewProducts creates a new products handler with a logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// getProducts returns the products from data sourse
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

// addProduct insert new product to the datastore
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
	rw.WriteHeader(http.StatusOK)
}

// updateProduct updates product by id
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

// MiddlewareProductValidation validates the product in the requests and calls next if ok
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// create new product object
		product := data.Product{}

		// deserialize JSON string to an object
		err := product.FromJSON(r.Body)

		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)

		// call the next handler, witch can be another middleware in the chain, or the final handler
		next.ServeHTTP(rw, req)
	})
}
