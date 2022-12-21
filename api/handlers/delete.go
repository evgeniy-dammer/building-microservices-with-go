package handlers

import (
	"github.com/evgeniy-dammer/building-microservices-with-go/api/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Removes product by id from the database
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// DeleteProduct deletes product by id
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	id := getProductID(r)

	p.l.Println("Handle DELETE Product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
