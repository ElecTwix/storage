package store

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrAuth     = errors.New("auth error")
)

type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
}
