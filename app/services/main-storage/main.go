package main

import (
	"github.com/St3plox/Gopher-storage/foundation/logger"
	"net/http"
)

//Must be remove in future

func main() {

	log := logger.New("Storage - API")
	if err := run(log); err != nil {
		log.Error().
			Str("startup")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
