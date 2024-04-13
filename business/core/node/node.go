package node

import "github.com/St3plox/Gopher-storage/foundation/storage"

const stringIdLength = 10

type Node struct {
	hashId     int
	Adress     string
	Port       string
	isAvaiable bool
}

func New(adress string, port string) (*Node, error) {

	randString := randStringRunes(stringIdLength)
	hashId, _, err := storage.Hash(randString, 1)
	if err != nil {
		return nil, err
	}

	return &Node{
		hashId:     hashId,
		Adress:     adress,
		Port:       port,
		isAvaiable: false,
	}, nil
}

func GenVirtual(node Node) *Node {
	//TODO: implement me
	return nil
}

func (n Node) HashID() int {
	return n.hashId
}

func (n *Node) CheckConnection() bool {
	//TODO: implement me
	return false
}

func Get(nodeID int, key string) (any, error) {
	return nil, nil
}

func Put(nodeID, key string, val any) error {
	return nil
}
