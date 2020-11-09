package handlers

import (
	"net/http"

	"github.com/nicholasjackson/building-microservices-youtube/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}


// swagger:route POST /products/deposit deposit depositCrypto
// deposit crypto
//
// responses:
//	200: depositResponse
//  422: errorValidation
//  501: errorResponse

// Deposit handles POST requests to deposit
func (p *Products) Deposit(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Input)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	//data.AddProduct(prod)
}
