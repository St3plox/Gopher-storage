package balance

import (
	"github.com/St3plox/Gopher-storage/foundation/linkedlist"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	"sync"
)

type HashSpace struct {
	nodes *linkedlist.LinkedList[RemoteStorer]
}

func NewHashSpace() *HashSpace {
	return &HashSpace{
		nodes: linkedlist.New[RemoteStorer](),
	}
}

func (hs *HashSpace) Get(key string) (any, int, error) {
	keyHash, _, err := storage.Hash(key, 1)
	if err != nil {
		return nil, 500, err
	}

	listNode := hs.nodes.FindClosestNode(keyHash)
	nodes := []RemoteStorer{*listNode.Val, *listNode.Next.Val, *listNode.Next.Next.Val}

	var wg sync.WaitGroup
	var getError error
	wg.Add(len(nodes))


	var val any
	var exist bool

	for _, n := range nodes {
		func(node RemoteStorer) {
			defer wg.Done()

			if node.CheckConnection() != true {
				//TODO: add failure handling
			}

			if val, exist, err = node.Get(key); err != nil {
				getError = err
			}

		}(n)
		
		if getError == nil{
			break
		}
	}

	if getError != nil {
		return nil, 500, err
	}

	code := 200
	if !exist {
		code = 404
	}

	return val, code, nil
}

func (hs *HashSpace) Put(key string, value any) error {
	keyHash, _, err := storage.Hash(key, 1)
	if err != nil {
		return err
	}

	listNode := hs.nodes.FindClosestNode(keyHash)
	nodes := []RemoteStorer{*listNode.Val, *listNode.Next.Val, *listNode.Next.Next.Val}

	var wg sync.WaitGroup
	wg.Add(len(nodes))
	var m sync.Mutex
	var putErr error

	for _, n := range nodes {
		go func(node RemoteStorer) {
			defer wg.Done()

			if n.CheckConnection() != true {
				return
			}

			if err := node.Put(key, value); err != nil {
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
func (hs *HashSpace) InitializeNodes(nodes []RemoteStorer) {
	for _, n := range nodes {
		hs.nodes.Insert(n.HashID(), n)
		n.CheckConnection()
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

	//TODO: finish implementation
	return nil
}
