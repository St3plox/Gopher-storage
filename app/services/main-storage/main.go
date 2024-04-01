package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/St3plox/Gopher-storage/app/services/main-storage/handlers"
	"github.com/St3plox/Gopher-storage/business/web/debug"
	"github.com/St3plox/Gopher-storage/foundation/logger"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	"github.com/ardanlabs/conf/v3"
	"github.com/rs/zerolog"
	defaultLog "log"
	"net/http"
	"os"
	"runtime"
	"time"

	_ "expvar"
)

var build = "develop"

func main() {
	log := logger.New("STORAGE - SERVICE")

	if err := run(log); err != nil {
		log.Error().Err(err).Msg("startup")
		os.Exit(1)
	}
}

func run(log *zerolog.Logger) error {

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.Info().Str("startup", "GOMAXPROCS").Int("GOMAXPROCS", runtime.
		GOMAXPROCS(0)).
		Str("BUILD", build).
		Msg("startup")

	// -------------------------------------------------------------------------
	// Configuration

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:10s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
			ShutdownTimeout time.Duration `conf:"default:20s,mask"`
			APIHost         string        `conf:"default::3000"`
			DebugHost       string        `conf:"default::4000"`
		}
		Storage struct {
			StoragePath     string `conf:"default:/var/lib/gopher-st"`
			PartitionNumber int    `conf:"default:10"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	const prefix = "STORAGE"
	help, err := conf.Parse(prefix, &cfg)

	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// -------------------------------------------------------------------------
	// App Starting

	log.Info().Str("version", build).Msg("starting service")
	defer log.Info().Msg("shotdown complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Info().Str("config", out).Msg("startup")

	// -------------------------------------------------------------------------
	// Storage

	st := storage.NewStorage(cfg.Storage.StoragePath, cfg.Storage.PartitionNumber)

	// -------------------------------------------------------------------------
	// Start Debug Service

	log.Info().
		Str("status", "debug v1 router started").
		Str("host", cfg.Web.DebugHost).
		Msg("startup")

	go func() {
		if err := http.ListenAndServe(cfg.Web.DebugHost, debug.StandardLibraryMux()); err != nil {
			log.Error().
				Str("status", "debug v1 router closed").
				Str("host", cfg.Web.DebugHost).
				Err(err).
				Msg("shutdown")

		}
	}()

	// -------------------------------------------------------------------------
	// Start API Service

	log.Info().Msg("initializing V1 API support")
	shutdown := make(chan os.Signal, 1)
	apiMux := handlers.APIMux(handlers.APIMuxConfig{
		Shutdown: shutdown,
		Log:      log,
		Storage:  st,
	})

	errorLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      apiMux,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     defaultLog.New(&errorLogger, "", 0),
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Info().
			Str("status", "api router started").
			Str("host", api.Addr).
			Msg("startup")
		serverErrors <- api.ListenAndServe()
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Info().
			Str("status", "shutdown started").
			Str("signal", sig.String()).
			Msg("shutdown")
		defer log.Info().Str("status", "shutdown complete").
			Str("signal", sig.String()).
			Msg("shutdown")

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			_ = api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}

	}

	return nil
}
