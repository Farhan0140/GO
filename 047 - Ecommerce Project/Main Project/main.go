package main

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func main() {
	cnf := config.GetConfig()

	middlewares := middlewares.NewMiddlewares(cnf)
	productRepo := repo.NewProductRepo()

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
