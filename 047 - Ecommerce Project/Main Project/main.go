package main

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
)

func main() {
	cnf := config.GetConfig()

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
