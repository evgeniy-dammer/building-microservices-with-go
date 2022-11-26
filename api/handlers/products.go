package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/evgeniy-dammer/building-microservices-with-go/api/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a new products handler with a logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point of the handler
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// check request method
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.updateProduct(rw, r)
		return
	}

	// if no method is satisfied
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from data sourse
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
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
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	// create new product object
	product := &data.Product{}

	// deserialize JSON string to an object
	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to decode JSON", http.StatusInternalServerError)
		return
	}

	// insert an object to the datastore
	err = data.InsertProduct(product)

	if err != nil {
		http.Error(rw, "Unable to add product", http.StatusInternalServerError)
		return
	}

	// if all is OK
	rw.WriteHeader(http.StatusOK)
}

// updateProduct updates product by id
func (p *Products) updateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")

	// creating regexp to find id in URL
	reg := regexp.MustCompile("/([0-9]+)")
	group := reg.FindAllStringSubmatch(r.URL.Path, -1)

	// if URL is invalid
	if len(group[0]) != 2 || len(group) != 1 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	// convert id to string format
	idString := group[0][1]
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	// create new product object
	product := &data.Product{}

	// deserialize JSON string to an object
	err = product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to decode JSON", http.StatusInternalServerError)
		return
	}

	// updates an object to the datastore
	err = data.UpdateProduct(id, product)

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
