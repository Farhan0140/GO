package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Nadim. I'm 24 years old....")
}

func main () {
	mux := http.NewServeMux()	// Router

	mux.HandleFunc("/hello", HelloHandler)	// Route 
	mux.HandleFunc("/about", AboutHandler)	// Route 

	fmt.Println("Server Running on port: 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("***Error Occurred", err)
	}
}