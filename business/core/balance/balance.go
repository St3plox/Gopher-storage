// Package balancer is used for balancing request in hash space
package balance

import (
	"fmt"
	"github.com/rs/zerolog"
)

type Balancer interface {
	Get(key string) (any, int, error)
	Put(key string, value any) error
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

func (c *Core) Get(key string) (any, int, error) {

	val, code, err := c.balancer.Get(key)
	if err != nil {
		c.log.Err(err).Msg("Error occured in get core")
		return nil, 500, fmt.Errorf("get: %w", err)
	}

	return val, code, nil

}

func (c *Core) Post(key string, value any) error {

	if err := c.balancer.Put(key, value); err != nil {
		c.log.Err(err).Msg("Error occured in post core")
		return fmt.Errorf("put: %w", err)
	}

	return nil
}
