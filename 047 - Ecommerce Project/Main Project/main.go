package main

import (
	"ecommerce/cmd"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func main() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	
	mux := http.NewServeMux() // Router
	wrappedMux := manager.WrapMux(mux)
	
	cmd.InitRoutes(mux, manager)
	fmt.Println("Server Running on port: 8080")

	err := http.ListenAndServe(":8080", wrappedMux)

	if err != nil {
		fmt.Println("***Error Occurred", err)
	}
}
