package main

import (
	"errors"
	"fmt"
	"github.com/St3plox/Gopher-storage/app/services/gateway-api/handlers"
	"github.com/St3plox/Gopher-storage/business/core/balance"
	"github.com/St3plox/Gopher-storage/business/core/node"
	"github.com/St3plox/Gopher-storage/business/web/v1/debug"
	"github.com/St3plox/Gopher-storage/foundation/logger"
	"github.com/ardanlabs/conf/v3"
	"github.com/rs/zerolog"

	defaultLog "log"

	"net/http"
	"os"
	"runtime"
	"time"
)

type Node struct {
	Address string `conf:"default:localhost"`
	Port    string `conf:"default:3000"`
}

var build = "develop"

const nodesCount = 20

func main() {
	log := logger.New("GATEWAY - SERVICE")

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
			APIHost         string        `conf:"default::8081"`
			DebugHost       string        `conf:"default::4001"`
		}
		Nodes []Node `conf:"default:0.0.0.0:3000"`
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	cfg.Nodes = make([]Node, 1, nodesCount)
	cfg.Nodes[0] = Node{Address: "0.0.0.0", Port: "3000"}

	const prefix = "GATEWAY"
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
	defer log.Info().Msg("shutdown complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Info().Str("config", out).Msg("startup")

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
	// Initializing Hash Space
	hs := balance.NewHashSpace()
	nodes := make([]balance.RemoteStorer, 0, nodesCount*3)

	for _, n := range cfg.Nodes {

		nd, err := node.New(n.Address, n.Port)
		if err != nil {
			panic(err)
		}

		virtN1, err := node.GenVirtual(nd)
		virtN2, err := node.GenVirtual(nd)
		if err != nil {
			panic(err)
		}

		nodes = append(nodes, nd)
		nodes = append(nodes, virtN1)
		nodes = append(nodes, virtN2)
	}

	hs.InitializeNodes(nodes)

	// -------------------------------------------------------------------------
	// Start API Service

	log.Info().Msg("initializing V1 API support")
	shutdown := make(chan os.Signal, 1)
	apiMux := handlers.APIMux(handlers.APIMuxConfig{
		Shutdown: shutdown,
		Log:      log,
		Balancer: hs,
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
	}
}
