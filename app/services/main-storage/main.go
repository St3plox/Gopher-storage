package main

import (
	"fmt"
	"github.com/St3plox/Gopher-storage/foundation/logger"
	"github.com/rs/zerolog"
	"os"
	"runtime"
	"github.com/ardanlabs/conf/v3"
)

var build = "develop"

type Config struct {
	APIHost   string `default:"0.0.0.0:3000"`
	DebugHost string `default:"0.0.0.0:4000"`
}

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
			APIHost   string `conf:"default:0.0.0.0:3000"`
			DebugHost string `conf:"default:0.0.0.0:4000"`
		}
		DB struct {
			StoragePath string `conf:"default:/var/lib/gopher-storage"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	// Print out the configuration
	fmt.Println("API Host:", cfg.Web.APIHost)
	fmt.Println("Debug Host:", cfg.DB.StoragePath)

	return nil
}
