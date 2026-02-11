package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middlewares"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /test",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	mux.Handle(
		"GET /products/{productID}",
		manager.With(
			http.HandlerFunc(handlers.GetProductById),
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)
}