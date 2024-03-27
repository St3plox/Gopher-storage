//Package storage used for handling saving/retreiving data on a single node
//server
package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Storage struct {
	DefaultPath     string
	PartitionNumber int
}

func NewStorage() *Storage {
	return &Storage{
		DefaultPath:     "/var/lib/gopher-storage",
		PartitionNumber: 10,
	}
}

func (s *Storage) Put(key string, value any) error {
	
	keyHash, partition, err := s.hash(key)
	if err != nil {
		return err
	}
	
	doc := NewDocument(key, value)
	jsonDoc, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	
	// Construct file path
	partitionDir := filepath.Join(s.DefaultPath, fmt.Sprintf("partition_%d", partition))
	filePath := filepath.Join(partitionDir, fmt.Sprintf("%d.json", keyHash))

	// Ensure partition directory exists
	if _, err := os.Stat(partitionDir); os.IsNotExist(err) {
		if err := os.MkdirAll(partitionDir, 0755); err != nil {
			return err
		}
	}

	// Write JSON document to file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonDoc)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Get(key string) (bool, error) {
	//TODO: impement me
	return false, nil
}
