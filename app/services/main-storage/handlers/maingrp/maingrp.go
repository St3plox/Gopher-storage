package maingrp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	"github.com/St3plox/Gopher-storage/foundation/web"
	"net/http"
)

//TODO: remove
var s *storage.Storage

func init() {
	s = storage.NewStorage()
}

// SaveData temporarily here
type SaveData struct {
	Key string
	Val any
}

func Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	key := r.PathValue("key")
	
	if key == "" {
		http.Error(w, "key cannot be empty", http.StatusBadRequest)
		return nil
	}
	
	val, exists, err := s.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return nil
	}

	response := struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}{
		Key:   key,
		Value: val,
	}

	err = web.Respond(ctx, w, response, http.StatusCreated)
	if err != nil {
		return err
	}
	return nil
}

func Post(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var sd SaveData
	err := json.NewDecoder(r.Body).Decode(&sd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	fmt.Fprintf(w, "Save data: %+v", sd)

	err = s.Put(sd.Key, sd.Val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return nil
}
