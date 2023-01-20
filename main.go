package main

import (
	"go-microservices/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	http.HandleFunc()
	http.ListenAndServe(":9090", nil)
}
