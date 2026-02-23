package product

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /test",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProductById),
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
		),
	)
}
