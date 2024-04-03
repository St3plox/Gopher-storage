package storage

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

const tempDir = "/var/lib/gopher-st-test/"
const partitionNumber = 10

func TestStorage_Put(t *testing.T) {

	err := generateTestDirs()
	if err != nil {
		t.Errorf("Error creating test dirs %v", err)
	}

	s := NewStorage(tempDir, partitionNumber)

	// Define your test cases
	var tests = []struct {
		name     string
		key      string
		value    any
		expected error
	}{
		{"put-single-string-test", "key0", "val", nil},
		{"put-struct-test", "key1", struct {
			name string
			age  uint8
		}{"Oleg", 23}, nil},
		{"same-key-test", "key0", "val", nil},
	}

	t.Log("Running tests for saving different data under the keys")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Perform the Put operation
			err := s.Put(test.key, test.value)
			if !errors.Is(err, test.expected) {
				t.Errorf("Put(%s, %s) returned unexpected error: %v", test.key, test.value, err)
			}
		})
	}
}

func TestStorage_Get(t *testing.T) {

	err := generateTestDirs()
	if err != nil {
		t.Errorf("Error creating test dirs %v", err)
	}

	s := NewStorage(tempDir, partitionNumber)

	key, val := "key", "val"
	err = s.Put(key, val)
	if err != nil {
		t.Errorf("Error saving key %v", err)
	}

	type fields struct {
		DefaultPath     string
		PartitionNumber int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		val     any
		isExist bool
		wantErr error
	}{
		{"exist-test", args{key}, val, true, nil},
		{"non-exist-test", args{"non-key"}, nil, false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := s.Get(tt.args.key)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.val) {
				t.Errorf("Get() got = %v, want %v", got, tt.val)
			}
			if got1 != tt.isExist {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.isExist)
			}
		})
	}
}

func generateTestDirs() error {
	err := os.MkdirAll(tempDir, 0755)
	if err != nil {
		return err
	}
	defer func() error {
		err = os.RemoveAll(tempDir)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		return err
	}

	return nil
}
