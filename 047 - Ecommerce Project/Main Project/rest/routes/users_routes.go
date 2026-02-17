package routs

import (
	"ecommerce/rest/handlers/user_handlers"
	"ecommerce/rest/middlewares"
	"net/http"
)

func UserRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
		),
	)

	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)
}