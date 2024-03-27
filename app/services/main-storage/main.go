package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /storage/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})
	
	mux.HandleFunc("PUT /storage", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err) 
	}
}