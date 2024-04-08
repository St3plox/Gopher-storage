package maingrp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	v1 "github.com/St3plox/Gopher-storage/business/web/v1"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	"github.com/St3plox/Gopher-storage/foundation/web"
	"net/http"
)

type Handler struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

// SaveData temporarily here
type SaveData struct {
	Key string
	Val any
}

func (h *Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	key := r.PathValue("key")

	if key == "" {
		http.Error(w, "key cannot be empty", http.StatusBadRequest)
		return nil
	}

	//TODO: fix error message not displayed in json
	val, exists, err := h.storage.Get(key)
	if err != nil {
		return v1.NewRequestError(errors.New("Sorage error"+err.Error()), http.StatusInternalServerError)
	}
	if !exists {
		return v1.NewRequestError(errors.New("KEY NOT FOUND"), http.StatusNotFound)
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

func (h *Handler) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var sd SaveData
	err := json.NewDecoder(r.Body).Decode(&sd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	fmt.Fprintf(w, "Save data: %+v", sd)

	err = h.storage.Put(sd.Key, sd.Val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	err = web.Respond(ctx, w, nil, http.StatusOK)
	if err != nil {
		return err
	}

	return nil
}
