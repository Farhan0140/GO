package util

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handleAllRequest := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}

		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handleAllRequest)
}
