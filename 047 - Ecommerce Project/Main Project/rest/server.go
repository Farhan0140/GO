package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux() // Router
	wrappedMux := manager.WrapMux(mux)

	InitRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server Running on port ", addr)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("***Error Occurred", err)
		os.Exit(1)
	}
}