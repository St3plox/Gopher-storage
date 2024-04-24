package gatewaygrp

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/St3plox/Gopher-storage/business/core/balance"
	v1 "github.com/St3plox/Gopher-storage/business/web/v1"
	"github.com/St3plox/Gopher-storage/foundation/web"
	"net/http"
)

type Handler struct {
	core *balance.Core
}

func New(core *balance.Core) *Handler {
	return &Handler{core: core}
}

type SaveData struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

func (h *Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	key := r.PathValue("key")

	if key == "" {
		http.Error(w, "key cannot be empty", http.StatusBadRequest)
		return nil
	}

	val, err := h.core.Get(key)

	if err != nil {
		return v1.NewRequestError(errors.New("Core error"+err.Error()), http.StatusInternalServerError)
	}

	//Refactor for status code
	/*
		if !exists {
			return v1.NewRequestError(errors.New("KEY NOT FOUND"), http.StatusNotFound)
		}
	*/
	response := struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}{
		Key:   key,
		Value: val,
	}

	err = web.Respond(ctx, w, response, http.StatusOK)
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

	err = h.core.Post(sd.Key, sd.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	err = web.Respond(ctx, w, nil, http.StatusCreated)
	if err != nil {
		return err
	}

	return nil
}
