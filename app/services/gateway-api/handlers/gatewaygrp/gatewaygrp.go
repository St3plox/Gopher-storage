package gatewaygrp

import "github.com/St3plox/Gopher-storage/business/core/balance"

type Handler struct {
	core *balance.Core
}

func New(core *balance.Core) *Handler{
	return &Handler{core: core}
}
