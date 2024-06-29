package memory

import (
	"context"
	"fmt"
)

// Store is a simple key-value store for embeddings.
type Store struct {
	store map[string][]float64
}

// NewStore creates a new Store from a redis client.
func NewStore() *Store {
	return &Store{store: make(map[string][]float64)}
}

// Get gets a value from the
func (s *Store) Get(
	ctx context.Context,
	utterance string,
) (embedding []float64, err error) {
	embedding, ok := s.store[utterance]
	if !ok {
		return nil, fmt.Errorf("key does not exist: %w", err)
	}
	return embedding, nil
}

// Set sets a value in the store
func (s *Store) Set(
	ctx context.Context,
	utterance string,
	value []float64,
) (string, error) {
	s.store[utterance] = value
	return utterance, nil
}
