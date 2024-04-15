package maingrp

import (
	"context"
	node_api "github.com/St3plox/Gopher-storage/business/proto/github.com/St3plox/Gopher-storage/app/services/node-api"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	anypb "github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	node_api.NodeV1Server
	storer storage.Storer
}

func New(storer storage.Storer) *Handler {
	return &Handler{
		storer: storer,
	}
}

func (h *Handler) Create(ctx context.Context, req *node_api.CreateRequest) (*emptypb.Empty, error) {

	err := h.storer.Put(req.Key, req.Val)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) Get(ctx context.Context, req *node_api.GetRequest) (*node_api.GetResponse, error) {

	val, _, err := h.storer.Get(req.Key)
	if err != nil {
		return nil, err
	}

	valPtr := &anypb.Any{}
	if val != nil {
		valPtr = val.(*anypb.Any)
	}

	//TODO: add bool value to the response
	resp := &node_api.GetResponse{Val: valPtr}

	return resp, nil
}

/*
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

	val, exists, err := h.storer.Get(key)
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

	err = h.storer.Put(sd.Key, sd.Val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	err = web.Respond(ctx, w, nil, http.StatusOK)
	if err != nil {
		return err
	}

	return nil
}*/
