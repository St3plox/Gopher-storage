package storage

import (
	"hash/fnv"
)

// Hash  returns hash value and partition of a string key
func (s *Storage) Hash(key string) (hash int, partition int, err error) {
	// Create a new FNV-1a hash
	h := fnv.New32a()

	// Write the key bytes to the hash
	_, err = h.Write([]byte(key))
	if err != nil {
		return 0, 0, err
	}

	// Get the hash value as a uint32
	hash = int(h.Sum32())
	partition = hash % s.PartitionNumber

	return hash, partition, nil
}
