package valkey

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/conneroisu/go-semantic-router/domain"
	"github.com/redis/go-redis/v9"
)

// Store is a simple key-value store for embeddings stored in valkey.
type Store struct {
	rds *redis.Client
}

// NewStore creates a new Store from a redis client.
func NewStore(rds *redis.Client) *Store {
	return &Store{rds: rds}
}

// Close closes the redis connection of the valkey store.
func (s *Store) Close() error {
	return s.rds.Close()
}

// Get gets a value from the valkey store.
func (s *Store) Get(
	ctx context.Context,
	utterance string,
) (embedding []float64, err error) {
	cmd := s.rds.Get(ctx, utterance)
	val, err := cmd.Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Println("key2 does not exist")
			fmt.Println(err)
			return nil, fmt.Errorf("key does not exist: %w", err)
		}
		return nil, err
	}
	var utPr domain.Utterance
	err = json.Unmarshal(bytes.NewBufferString(val).Bytes(), &utPr)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling embedding: %w", err)
	}
	return utPr.Embed, nil
}

// Set sets a value in the valkey store.
func (s *Store) Store(
	ctx context.Context,
	utterance domain.Utterance,
) error {
	val, err := json.Marshal(utterance)
	if err != nil {
		return fmt.Errorf("error marshaling embedding: %w", err)
	}
	cmd := s.rds.Set(
		ctx,
		utterance.Utterance,
		string(val),
		0,
	)
	err = cmd.Err()
	if err != nil {
		return fmt.Errorf("error setting embedding: %w", err)
	}
	return nil
}
