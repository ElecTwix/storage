package local

import (
	"os"
	"path"
	"sync"
	"time"

	"github.com/ElecTwix/cache"
	"github.com/ElecTwix/storage/pkg/store"
)

type LocalStore struct {
	mutex         sync.Mutex
	cache         *cache.Cache[string, []byte]
	cacheDuration time.Duration
	Path          string
}

func NewLocalStore(path string, tickTime, timeout time.Duration) *LocalStore {
	ticker := time.NewTicker(tickTime)
	LocalStoreData := &LocalStore{
		cache:         cache.NewCache[string, []byte](*ticker),
		cacheDuration: timeout,
		Path:          path,
	}
	return LocalStoreData
}

func (s *LocalStore) Get(key string) ([]byte, error) {
	v, ok := s.cache.Get(key)
	if ok {
		return v, nil
	}

	data, err := os.ReadFile(path.Join(s.Path, key))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, store.ErrNotFound
		}
		return nil, err
	}
	return data, err
}

func (s *LocalStore) Set(key string, data []byte) error {
	s.mutex.Lock()
	s.cache.Set(key, data, s.cacheDuration)
	defer s.mutex.Unlock()

	return os.WriteFile(path.Join(s.Path, key), data, 0644)
}
