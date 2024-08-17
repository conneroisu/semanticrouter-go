package memory

import (
	"context"
	"fmt"

	semanticrouter "github.com/conneroisu/go-semantic-router"
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
	_ context.Context,
	utterance string,
) (embedding []float64, err error) {
	embedding, ok := s.store[utterance]
	if !ok {
		return nil, fmt.Errorf("key does not exist: %w", err)
	}
	return embedding, nil
}

// Store sets a value in the store
func (s *Store) Store(
	_ context.Context,
	utterance semanticrouter.Utterance,
) error {
	s.store[utterance.Utterance] = utterance.Embed
	return nil
}

// Close closes the store.
func (s Store) Close() error {
	return nil
}
