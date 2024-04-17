package handlers

import (
	"github.com/St3plox/Gopher-storage/app/services/node-api/handlers/maingrp"
	"github.com/St3plox/Gopher-storage/business/web/v1/mid"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	"github.com/St3plox/Gopher-storage/foundation/web"
	"github.com/rs/zerolog"
	"os"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zerolog.Logger
	Storer   storage.Storer
}

func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())

	h := maingrp.New(cfg.Storer)

	app.Handle("GET /storage/{key}", h.Get)
	app.Handle("POST /storage", h.Post)

	return app
}

/*
type GRPCServerConfig struct {
	Shutdown chan os.Signal
	Log      *zerolog.Logger
	Storer   storage.Storer
}

func GRPCServer(cfg GRPCServerConfig) *grpc.Server {

	server := grpc.NewServer()
	h := maingrp.New(cfg.Storer)

	node_api.RegisterNodeV1Server(server, h)

	return server
}*/
