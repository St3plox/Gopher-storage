package main

import (
	"encoding/json"
	"fmt"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	"log"
	"net/http"
)

//Must be remove in future
var s *storage.Storage

func init() {
	s = storage.NewStorage()
}

// SaveData temporarily here
type SaveData struct {
	Key string
	Val any
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /storage/{key}", func(w http.ResponseWriter, r *http.Request) {

		key := r.PathValue("key")
		if key == "" {
			http.Error(w, "key cannot be empty", http.StatusBadRequest)
			return
		}
		log.Default().Printf("--- Got key %s --- ", key)

		val, exists, err := s.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !exists {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}

		response := struct {
			Key   string      `json:"key"`
			Value interface{} `json:"value"`
		}{
			Key:   key,
			Value: val,
		}


		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}

	})

	//TODO: fix put errors
	mux.HandleFunc("POST /storage", func(w http.ResponseWriter, r *http.Request) {

		var sd SaveData
		err := json.NewDecoder(r.Body).Decode(&sd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Default().Println("-------------------- Successfuly decoded response body --------------------")

		fmt.Fprintf(w, "Save data: %+v", sd)

		err = s.Put(sd.Key, sd.Val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Default().Println("-------------------- Successfuly saved data --------------------")
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
