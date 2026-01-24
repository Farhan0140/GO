package main

import (
	"ecommerce/handlers"
	"ecommerce/util"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // Router

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))                // Route
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))             // Route
	mux.Handle("GET /products/{productID}", http.HandlerFunc(handlers.GetProductById)) // Route

	fmt.Println("Server Running on port: 8080")

	err := http.ListenAndServe(":8080", util.GlobalRouter(mux))

	if err != nil {
		fmt.Println("***Error Occurred", err)
	}
}
