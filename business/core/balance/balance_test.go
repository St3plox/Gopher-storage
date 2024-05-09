package balance

import (
	"github.com/St3plox/Gopher-storage/business/core/node"
	"testing"
)

type mockNode struct {
	*node.Node
	GetFunc func(key string) (interface{}, int, error)
	PutFunc func(key string, value interface{}) (int, error)
}

func (m *mockNode) Get(key string) (interface{}, int, error) {
	if m.GetFunc != nil {
		return m.GetFunc(key)
	}
	return nil, 0, nil // Default behavior if GetFunc is not set
}

func (m *mockNode) Put(key string, value interface{}) (int, error) {
	if m.PutFunc != nil {
		return m.PutFunc(key, value)
	}
	return 0, nil // Default behavior if PutFunc is not set
}

func TestHashSpace_Get(t *testing.T) {
	// Mocking the node and its behavior
	mock := &mockNode{
		Node: &node.Node{}, // Embedding the node.Node type
		GetFunc: func(key string) (interface{}, int, error) {
			return "test_value", 200, nil
		},
	}

	hs := NewHashSpace()
	hs.InitializeNodes([]mockNode{mock}) // Casting mock to node.Node

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
		PutFunc: func(key string, value interface{}) (int, error) {
			return 201, nil
		},
	}

	hs := NewHashSpace()
	hs.InitializeNodes([]node.Node{mock}) // Casting mock to node.Node

	err := hs.Put("test_key", "test_value")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
