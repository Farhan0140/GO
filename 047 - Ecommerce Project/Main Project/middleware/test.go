package middleware

import (
	"log"
	"net/http"
)

func Test(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Fuck You Middleware")

		next.ServeHTTP(w, r)
	})
}