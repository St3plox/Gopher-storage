package node

import "github.com/St3plox/Gopher-storage/foundation/storage"

const stringIdLength = 10

type Node struct {
	hashId int
	Adress string
	Port   string
}

func NewNode(adress string, port string) (*Node, error) {

	randString := randStringRunes(stringIdLength)
	hashId, _, err := storage.Hash(randString, 1)
	if err != nil {
		return nil, err
	}

	return &Node{
		hashId: hashId,
		Adress: adress,
		Port:   port,
	}, nil
}
