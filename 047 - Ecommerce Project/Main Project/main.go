package main

import (
	"ecommerce/middleware"
	"ecommerce/util"
	"ecommerce/cmd"
	"fmt"
	"net/http"
)

func main() {
	manager := middleware.NewManager()
	manager.Use(middleware.Test, middleware.Logger)

	mux := http.NewServeMux() // Router

	cmd.InitRoutes(mux, manager)

	fmt.Println("Server Running on port: 8080")

	err := http.ListenAndServe(":8080", util.GlobalRouter(mux))

	if err != nil {
		fmt.Println("***Error Occurred", err)
	}
}
