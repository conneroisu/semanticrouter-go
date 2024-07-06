package memory

import (
	"context"
	"fmt"

	"github.com/conneroisu/go-semantic-router/domain"
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
func (s Store) Get(
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
func (s Store) Store(
	_ context.Context,
	utterance domain.Utterance,
) error {
	var err error
	s.store[utterance.Utterance], err = utterance.Embedding()
	if err != nil {
		return fmt.Errorf("error getting embedding: %w", err)
	}
	return nil
}
