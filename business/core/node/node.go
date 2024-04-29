package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	"net/http"
)

const stringIdLength = 10

type Node struct {
	hashId      int
	Adress      string
	Port        string
	IsAvailable bool
	IsVirtual   bool
}

func New(adress string, port string) (*Node, error) {

	randString := randStringRunes(stringIdLength)
	hashId, _, err := storage.Hash(randString, 1)
	if err != nil {
		return nil, err
	}

	return &Node{
		hashId:      hashId,
		Adress:      adress,
		Port:        port,
		IsAvailable: false,
		IsVirtual:   false,
	}, nil
}

func GenVirtual(node *Node) (*Node, error) {

	virtualNode, err := New(node.Adress, node.Port)
	if err != nil {
		return nil, err
	}
	virtualNode.IsAvailable = node.IsAvailable
	virtualNode.IsVirtual = true

	return virtualNode, nil
}

// HashID Getter node's hash id
func (n *Node) HashID() int {
	return n.hashId
}

// CheckConnection function used to check if the node is avaiable
func (n *Node) CheckConnection() bool {

	addr := fmt.Sprintf("http://%s:%s/liveness", n.Adress, n.Port)

	resp, err := http.Get(addr)
	if err != nil  || resp.StatusCode != http.StatusOK{
		return false
	}
	defer resp.Body.Close()

	n.IsAvailable = true
	return true
}

// Get function sends Get request to this node address, returns value, response code, error
func (n *Node) Get(key string) (any, int, error) {

	addr := fmt.Sprintf("http://%s:%s/storage/%s", n.Adress, n.Port, key)

	resp, err := http.Get(addr)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer resp.Body.Close()

	var respBody storage.SaveData
	if err = json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return respBody, resp.StatusCode, nil
}

// Put function runs  Post request to this node, returns status code and error
func (n *Node) Put(key string, val any) (int, error) {

	post := struct {
		Key   string
		Value any
	}{Key: key, Value: val}

	jsonVal, err := json.Marshal(&post)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	addr := fmt.Sprintf("http://%s:%s/storage", n.Adress, n.Port)
	fmt.Println(addr)

	resp, err := http.Post(addr, "aplication/json", bytes.NewBuffer(jsonVal))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
