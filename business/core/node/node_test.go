package node

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNode_CheckConnection(t *testing.T) {
	// Mocking the HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/liveness" {
			http.Error(w, "Invalid path", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	node, err := New("127.0.0.1", ts.URL[len(ts.URL)-5:])

	if err != nil {
		t.Error(err)
	}

	// Testing connection check
	if !node.CheckConnection() {
		t.Error("Expected connection to be available, got false")
	}
}

func TestNode_Get(t *testing.T) {
	// Mocking the HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/storage/test_key" {
			http.Error(w, "Invalid path", http.StatusBadRequest)
			return
		}

		// Simulate response body
		respBody := `{"data": "test_data"}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(respBody))
	}))
	defer ts.Close()

	node, err := New("127.0.0.1", ts.URL[len(ts.URL)-5:])
	if err != nil {
		t.Error(err)
	}

	// Testing Get function
	data, statusCode, err := node.Get("test_key")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, statusCode)
	}
	if data == nil {
		t.Errorf("Expected non-nil response data, got nil")
	}
}

func TestNode_Put(t *testing.T) {
	// Mocking the HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/storage" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Simulate a successful response
		w.WriteHeader(http.StatusCreated)
	}))
	defer ts.Close()

	node, err := New("127.0.0.1", ts.URL[len(ts.URL)-5:])
	if err != nil {
		t.Error(err)
	}

	// Testing Put function
	statusCode, err := node.Put("test_key", "test_value")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if statusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, statusCode)
	}
}

func TestNode_GenVirtual(t *testing.T) {

	node, err := New("localhost", "8080")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	vNode, err := GenVirtual(node)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if vNode.IsVirtual != true {
		t.Errorf("Expected node to be virtual, got false")
	}

	if vNode.Adress != vNode.Adress && node.Port != node.Port {
		t.Errorf("Expected virtual node's addres and port to be equal, got virtual: %s:%s node: %s:%s", vNode.Adress, vNode.Port, node.Adress, node.Port)
	}
	
	if vNode.hashId == node.hashId {
		t.Errorf("Expected virtual node hash id to be different from node")
	}


}
