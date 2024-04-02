package mid

import (
	"context"
	v1 "github.com/St3plox/Gopher-storage/business/web/v1"
	"github.com/St3plox/Gopher-storage/foundation/web"
	"github.com/rs/zerolog"
	"net/http"
)

// Errors handles errors coming out of the call chain. It detects normal
// application errors which are used to respond to the client in a uniform way.
// Unexpected errors (status >= 500) are logged.
func Errors(log *zerolog.Logger) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if err := handler(ctx, w, r); err != nil {

				log.Err(err).Str("trace_id", web.GetTraceID(ctx)).Msg("message")

				er := v1.ErrorResponse{
					Error: http.StatusText(http.StatusInternalServerError),
				}
				status := http.StatusInternalServerError

				if err := web.Respond(ctx, w, er, status); err != nil {
					return err
				}

				// If we receive the shutdown err we need to return it
				// back to the base handler to shut down the service.
				if web.IsShutdown(err) {
					return err
				}
			}

			return nil
		}

		return h
	}

	return m
}
