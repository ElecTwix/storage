package storage

import "github.com/ElecTwix/storage/pkg/store"

type Storage struct {
	storeInterface store.Store
}

func NewStorage(storeInterface store.Store) *Storage {
	return &Storage{storeInterface: storeInterface}
}

func (s *Storage) Get(key string) ([]byte, error) {
	return s.storeInterface.Get(key)
}

func (s *Storage) Set(key string, data []byte) error {
	return s.storeInterface.Set(key, data)
}
