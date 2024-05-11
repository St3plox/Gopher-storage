package balance

import (
	"github.com/St3plox/Gopher-storage/business/core/node"
	"testing"
)

type mockNode struct {
	*node.Node
	GetFunc func(key string) (interface{}, bool, error)
	PutFunc func(key string, value interface{}) error
}

func (m *mockNode) Get(key string) (interface{}, bool, error) {
	if m.GetFunc != nil {
		return m.GetFunc(key)
	}
	return nil, false, nil // Default behavior if GetFunc is not set
}

func (m *mockNode) Put(key string, value interface{}) error {
	if m.PutFunc != nil {
		return m.PutFunc(key, value)
	}
	return nil // Default behavior if PutFunc is not set
}

func TestHashSpace_Get(t *testing.T) {
	// Mocking the node and its behavior
	mock := &mockNode{
		Node: &node.Node{},
		GetFunc: func(key string) (interface{}, bool, error) {
			return "test_value", true, nil
		},
	}

	hs := NewHashSpace()
	hs.InitializeNodes([]RemoteStorer{mock}) // Casting mock to node.Node

	value, statusCode, err := hs.Get("test_key")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "test_value" {
		t.Errorf("Expected value 'test_value', got %v", value)
	}
	if statusCode != 200 {
		t.Errorf("Expected status code 200, got %d", statusCode)
	}
}

func TestHashSpace_Put(t *testing.T) {
	// Mocking the node and its behavior
	mock := &mockNode{
		Node: &node.Node{}, // Embedding the node.Node type
		PutFunc: func(key string, value interface{}) error {
			return nil
		},
	}

	hs := NewHashSpace()
	hs.InitializeNodes([]RemoteStorer{mock}) // Casting mock to node.Node

	err := hs.Put("test_key", "test_value")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
