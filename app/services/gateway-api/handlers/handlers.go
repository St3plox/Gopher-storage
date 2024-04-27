package handlers

import (
	"github.com/St3plox/Gopher-storage/app/services/gateway-api/handlers/gatewaygrp"
	"github.com/St3plox/Gopher-storage/business/core/balance"
	"github.com/St3plox/Gopher-storage/business/web/v1/mid"
	"github.com/St3plox/Gopher-storage/foundation/web"
	"github.com/rs/zerolog"
	"os"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zerolog.Logger
	Balancer balance.Balancer
}

func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())

	core := balance.NewCore(cfg.Log, cfg.Balancer)
	h := gatewaygrp.New(core)

	app.Handle("GET /storage/{key}", h.Get)
	app.Handle("POST /storage", h.Post)
	return app
}
