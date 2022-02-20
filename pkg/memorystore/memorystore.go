// Package memorystore provides read and write access to memory.
package memorystore

import (
	"sync"
)

// MemoryStore represents a file in memory.
type MemoryStore struct {
	content string
	m       *sync.RWMutex
}

// New returns a local filesystem object with a file path.
func New() *MemoryStore {
	return &MemoryStore{
		m: &sync.RWMutex{},
	}
}

// Load returns a file contents from the filesystem.
func (s *MemoryStore) Load() ([]byte, error) {
	s.m.RLock()
	b := []byte(s.content)
	s.m.RUnlock()
	return b, nil
}

// Save writes a file to the filesystem and returns an error if one occurs.
func (s *MemoryStore) Save(b []byte) error {
	s.m.Lock()
	s.content = string(b)
	s.m.Unlock()
	return nil
}
