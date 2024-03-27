package storage

import "time"

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

func save()  {
	
}
