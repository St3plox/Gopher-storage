package main

import (
	"encoding/json"
	"fmt"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	"net/http"
)

//Must be remove in future
var s *storage.Storage

func init()  {
	s = storage.NewStorage()
}

// SaveData temporarily here
type SaveData struct {
	Key string
	Val any
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /storage/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})

	mux.HandleFunc("PUT /storage", func(w http.ResponseWriter, r *http.Request) {

		var sd SaveData
		err := json.NewDecoder(r.Body).Decode(&sd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Save data: %+v", sd)

		err = s.Put(sd.Key, sd.Val)
		//TODO: Fix the error 
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
