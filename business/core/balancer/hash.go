package balancer

import (
	"github.com/St3plox/Gopher-storage/foundation/linkedlist"
)

type HashSpace struct {
	partitionNumber uint
	nodes           *linkedlist.LinkedList
}

func NewHashSpace(papartitionNumber uint) *HashSpace{
	return &HashSpace{
		partitionNumber: papartitionNumber,
		nodes: linkedlist.New(),
	}
}
