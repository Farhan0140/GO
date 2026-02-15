package rest

import (
	"ecommerce/rest/middlewares"
	routs "ecommerce/rest/routes"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	routs.ProductRoutes(mux, manager)
}