package balance

import (
	"github.com/St3plox/Gopher-storage/business/core/node"
	"github.com/St3plox/Gopher-storage/foundation/linkedlist"
	"github.com/St3plox/Gopher-storage/foundation/storage"
)

type HashSpace struct {
	nodes *linkedlist.LinkedList[*node.Node]
}

func NewHashSpace() *HashSpace {
	return &HashSpace{
		nodes: linkedlist.New[*node.Node](),
	}
}

func (hs *HashSpace) Get(key string) (any, error) {
	keyHash, _, err := storage.Hash(key, 1)
	if err != nil {
		return nil, err
	}

	listNode := hs.nodes.FindClosestNode(keyHash)

	val, err := node.Get(listNode.HashedID(), key)
	if err != nil {
		return nil, err
	}
	
	return val, nil
}

func (hs *HashSpace) Put(key string, value any) error {
	
	keyHash, _, err := storage.Hash(key, 1)
	if err != nil {
		return err
	}

	listNode := hs.nodes.FindClosestNode(keyHash)
	//TODO: add avaiability check
	
	//TODO: implement me
	return nil
}

// InitializeNodes inserts nods in node array before establishing connection1
func (hs *HashSpace) InitializeNodes(nodes []node.Node) {
	for _, node := range nodes {
		hs.nodes.Insert(node.HashID(), &node)
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

func getClosestNodeID(){
	//TODO: implement me
}
