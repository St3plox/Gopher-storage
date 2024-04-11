// Package balancer is used for balancing request in hash space
package balancer

import "github.com/rs/zerolog"

type Balancer interface {
	Get(key string)
	Put(key string, value any)
}

type Core struct {
	log      *zerolog.Logger
	balancer Balancer
}

func NewCore(log *zerolog.Logger, balancer Balancer) *Core {
	return &Core{
		log:      log,
		balancer: balancer,
	}
}

