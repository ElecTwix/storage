package local_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/ElecTwix/storage/pkg/store/local"
)

func TestNewLocalStore(t *testing.T) {

	path := "temp"
	tickTime := time.Second
	timeout := time.Second * 10
	store := local.NewLocalStore(path, tickTime, timeout)

	// Get data that not exist
	// Should return error
	_, err := store.Get("key")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Set data
	// Should not return error

	value := []byte("value")

	err = store.Set("key", value)
	if err != nil {
		t.Error("Expected nil, got error")
	}

	// Get data that exist
	// Should not return error
	data, err := store.Get("key")
	if err != nil {
		t.Error("Expected nil, got error")
	}

	// Check data
	// Should be equal
	if !bytes.Equal(data, value) {
		t.Error("Expected equal, got not equal")
	}

}
