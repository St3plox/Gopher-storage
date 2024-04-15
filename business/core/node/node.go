package node

import (
	"context"
	nodeapi "github.com/St3plox/Gopher-storage/business/proto/github.com/St3plox/Gopher-storage/app/services/node-api"
	"github.com/St3plox/Gopher-storage/foundation/storage"
	anypb "github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

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

func GenVirtual(node Node) (*Node, error) {

	virtualNode, err := New(node.Adress, node.Port)
	if err != nil {
		return nil, err
	}
	virtualNode.isAvaiable = node.isAvaiable

	return virtualNode, nil
}

// HashID Getter node's hash id
func (n *Node) HashID() int {
	return n.hashId
}

// CheckConnection function used to check if the node is avaiable
func (n *Node) CheckConnection() bool {
	//TODO: implement me
	return false
}

// Get function sends gRPC Get request to this node address
func (n *Node) Get(key string) (any, error) {

	conn, err := grpc.Dial(n.Adress+":"+n.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := nodeapi.NewNodeV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Get(ctx, &nodeapi.GetRequest{Key: key})
	if err != nil {
		return nil, err
	}

	return resp.Val, nil
}

// Put function runs gRPC Create request to this node
func (n *Node) Put(key string, val any) error {

	valPtr := &anypb.Any{}
	if val != nil {
		valPtr = val.(*anypb.Any)
	}

	conn, err := grpc.Dial(n.Adress+":"+n.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	c := nodeapi.NewNodeV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Create(ctx, &nodeapi.CreateRequest{Key: key, Val: valPtr})
	if err != nil {
		return err
	}
	_ = resp

	return nil
}
