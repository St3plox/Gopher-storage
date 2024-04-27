package balance

import (
	"github.com/St3plox/Gopher-storage/business/core/node"
	"github.com/St3plox/Gopher-storage/foundation/linkedlist"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	"sync"
)

type HashSpace struct {
	nodes *linkedlist.LinkedList[*node.Node]
}

func NewHashSpace() *HashSpace {
	return &HashSpace{
		nodes: linkedlist.New[*node.Node](),
	}
}

func (hs *HashSpace) Get(key string) (any, int, error) {
	keyHash, _, err := storage.Hash(key, 1)
	if err != nil {
		return nil, 500, err
	}

	listNode := hs.nodes.FindClosestNode(keyHash)
	n := listNode.Val

	val, code, err := (*n).Get(key)
	if err != nil {
		return nil, 500, err
	}

	//TODO: Add failure handling

	return val, code, nil
}

func (hs *HashSpace) Put(key string, value any) error {
	keyHash, _, err := storage.Hash(key, 1)
	if err != nil {
		return err
	}

	listNode := hs.nodes.FindClosestNode(keyHash)
	nodes := []*node.Node{*listNode.Val, *listNode.Next.Val, *listNode.Next.Next.Val}

	var wg sync.WaitGroup
	wg.Add(len(nodes))
	var m sync.Mutex
	var putErr error

	for _, n := range nodes {
		go func(node *node.Node) {
			defer wg.Done()

			if _, err := node.Put(key, value); err != nil {
				m.Lock()
				putErr = err
				m.Unlock()
			}
		}(n)
	}

	wg.Wait()

	//TODO: add failure handling
	//TODO: add avaiability check

	return putErr
}

// InitializeNodes inserts nods in node array before establishing connection1
func (hs *HashSpace) InitializeNodes(nodes []node.Node) {
	for _, n := range nodes {
		hs.nodes.Insert(n.HashID(), &n)
	}
}

func (hs *HashSpace) EstablishConnection() error {

	nodes := hs.nodes.Values()
	unanavailableNodeIds := make([]int, len(nodes))
	for i := range nodes {

		n := nodes[i]
		isAvailable := n.CheckConnection()

		if !isAvailable {
			unanavailableNodeIds = append(unanavailableNodeIds, n.HashID())
		}
	}

	return nil
}
