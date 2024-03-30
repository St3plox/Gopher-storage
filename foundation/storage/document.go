package storage

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Document struct {
	Key       string
	Value     any
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewDocument(key string, val any) *Document {
	return &Document{
		Key:       key,
		Value:     val,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func readDocumentFromFile(filePath string) (*Document, error) {
	// Read file contents
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into a Document struct
	var doc Document
	err = json.Unmarshal(data, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}
