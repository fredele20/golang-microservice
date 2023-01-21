package handlers

import (
	"context"
	"fmt"
	"go-microservices/product-api/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	listProduct := data.GetProducts()
	err := listProduct.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	product := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&product)

}

func (p Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "unable to convert ID", http.StatusBadRequest)
	}
	p.l.Println("Handle PUT Product", id)
	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, &product)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct {}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := data.Product{}

		err := product.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validate the product
		err = product.Validate()
		if err != nil {
			p.l.Println("ERROR validating product", err)
			http.Error(w, fmt.Sprintf("ERROR validating the product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)
		
		next.ServeHTTP(w, r)

	})
}
