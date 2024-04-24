// Package storage used for handling saving/retreiving data on a single node
// server
package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	DefaultPartitionFormat = "partition_%d"
)

type Storer interface {
	Put(key string, value interface{}) error
	Get(key string) (any, bool, error)
}

type Storage struct {
	DefaultPath     string
	PartitionNumber int
}

func NewStorage(defaultPath string, partitionNumber int) *Storage {
	return &Storage{
		DefaultPath:     defaultPath,
		PartitionNumber: partitionNumber,
	}
}

// Put function is for saving value under key string
// By default saves under path/gopher-storage/partition_n/key_hash/number.json
func (s *Storage) Put(key string, value interface{}) error {
	// Compute hash and partition
	keyHash, partition, err := Hash(key, s.PartitionNumber)
	if err != nil {
		return err
	}

	// Construct directory path
	partitionDir := s.partitionDirGenerate(keyHash, partition)

	// Ensure partition directory exists
	if _, err := os.Stat(partitionDir); os.IsNotExist(err) {
		if err := os.MkdirAll(partitionDir, 0755); err != nil {
			return err
		}
	}

	_, fileIndex, _, err := handleCollision(partitionDir, key)
	if err != nil {
		return err
	}

	fileName := strconv.Itoa(fileIndex) + ".json"

	// Construct file path
	filePath := filepath.Join(partitionDir, fileName)

	// Create document
	doc := NewDocument(key, value)
	jsonDoc, err := json.Marshal(doc)
	if err != nil {
		return err
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

func (s *Storage) Get(key string) (any, bool, error) {
	keyHash, partition, err := Hash(key, s.PartitionNumber)
	if err != nil {
		return nil, false, err
	}

	partitionDir := s.partitionDirGenerate(keyHash, partition)
	doc, _, isExist, err := handleCollision(partitionDir, key)
	if err != nil {
		fmt.Println("soser1")
		return nil, false, err
	}

	// Check if the document exists
	if doc == nil {
		return nil, false, nil // Return nil value and false for existence flag
	}

	return doc.Value, isExist, nil
}

// func partitionDirGenerate Generates partition dir.
// By default saves under path/gopher-storage/partition_n/key_hash/number.json
func (s *Storage) partitionDirGenerate(keyHash int, patrtition int) string {
	return filepath.Join(s.DefaultPath, fmt.Sprintf(DefaultPartitionFormat, patrtition), fmt.Sprintf("%d", keyHash))
}

// Function handleCollision handles collision of the hash value.
// Returns the *Document if there is such key, number of a file that needs to be saved, and bool isKeyExist
// Takes as an argument filepath without "n.json" and unhashed key
func handleCollision(partitionDir string, key string) (*Document, int, bool, error) {
	// Initialize variables
	maxIndex := 0
	keyExists := false
	var foundDoc *Document

	// Check if the partition directory exists
	if _, err := os.Stat(partitionDir); os.IsNotExist(err) {
		// If the directory doesn't exist, return indicating the key doesn't exist
		return nil, 0, false, nil
	}

	// Iterate through all files in the partition directory
	err := filepath.Walk(partitionDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a regular file
		if info.Mode().IsRegular() {
			// Parse the index from the file name (excluding the ".json" extension)
			fileName := strings.TrimSuffix(info.Name(), ".json")
			fileIndex, err := strconv.Atoi(fileName)
			if err != nil {
				return err
			}

			// Update the maximum index if the current index is greater
			if fileIndex > maxIndex {
				maxIndex = fileIndex
			}

			// Read the document from the file to check if it corresponds to the key
			doc, err := readDocumentFromFile(path)
			if err != nil {
				return err
			}

			// If the key matches, set keyExists to true and store the document
			if doc.Key == key {
				keyExists = true
				foundDoc = doc
			}
		}

		return nil
	})

	// If an error occurred during the walk, return the error
	if err != nil {
		return nil, 0, false, err
	}

	// Return the document, maximum index, and keyExists
	return foundDoc, maxIndex, keyExists, nil
}
