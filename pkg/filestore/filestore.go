// Package filestore provides read and write access to local filesystem.
package filestore

import (
	"io/ioutil"
	"sync"
)

// FileStore represents a file on the filesytem.
type FileStore struct {
	path string
	m    *sync.RWMutex
}

// New returns a local filesystem object with a file path.
func New(path string) *FileStore {
	return &FileStore{
		path: path,
		m:    &sync.RWMutex{},
	}
}

// Load returns a file contents from the filesystem.
func (s *FileStore) Load() ([]byte, error) {
	s.m.RLock()
	b, err := ioutil.ReadFile(s.path)
	s.m.RUnlock()
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Save writes a file to the filesystem and returns an error if one occurs.
func (s *FileStore) Save(b []byte) error {
	s.m.Lock()
	err := ioutil.WriteFile(s.path, b, 0644)
	s.m.Unlock()
	if err != nil {
		return err
	}

	return nil
}
