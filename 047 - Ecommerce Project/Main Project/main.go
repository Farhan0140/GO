package main

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"os"
)

func main() {
	cnf := config.GetConfig()
	dbCon, err := db.NewConnection(cnf.DB)
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middlewares.NewMiddlewares(cnf)
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
