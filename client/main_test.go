package main

import (
	"fmt"
	"go-microservices/client/products"
	"testing"
)


func TestClient(t *testing.T) {
	cfg := DefaultTransportConfig().WithHost("localhost:9090")
	c := NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v",prod.GetPayload()[0])
	t.Fail()
}