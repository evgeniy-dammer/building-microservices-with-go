// Package classification of Product API
//
// Documentation for Product API
//
//		Schemes: http
//		BasePath: /
//		Version: 1.0.0
//
//		Consumes:
//		- application/json
//
//	 Produces:
//	 - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"github.com/evgeniy-dammer/building-microservices-with-go/api/data"
	"log"
	"net/http"
)

// productsResponseWrapper is a wrapper for a list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// productsRequestWrapper is a wrapper for a product in POST request
// swagger:parameters productsRequest
type productsRequestWrapper struct {
	// A product sending in a POST request
	// in: query
	Body data.Product
}

// productsNoContentWrapper is a wrapper for no content returns in the response
// swagger:response noContent
type productsNoContentWrapper struct {
	// no content inside
}

// productIDParameterWrapper is a wrapper for id parameter in path
// swagger:parameters deleteProduct updateProduct
type productIDParameterWrapper struct {
	// The id of the product to update or delete from database
	// in: path
	// required: true
	ID int `json:"id"`
}

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// KeyProduct
type KeyProduct struct{}

// NewProducts creates a new products handler with a logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
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

		// validate product
		err = product.Validate()

		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)

		// call the next handler, witch can be another middleware in the chain, or the final handler
		next.ServeHTTP(rw, req)
	})
}
