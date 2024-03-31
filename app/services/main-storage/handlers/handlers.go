package handlers

import (
	"github.com/St3plox/Gopher-storage/app/services/main-storage/handlers/maingrp"
	"github.com/St3plox/Gopher-storage/business/web/v1/mid"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	"github.com/St3plox/Gopher-storage/foundation/web"
	"github.com/rs/zerolog"
	"os"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zerolog.Logger
	Storage  *storage.Storage
}

func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log))
	
	app.Handle("GET /storage/{key}", maingrp.Get)
	app.Handle("POST /storage", maingrp.Post)
	
	return app
}