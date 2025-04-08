package valkey

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/conneroisu/semanticrouter-go"
	"github.com/redis/go-redis/v9"
)

// Store is a valkey/redis store for embeddings.
type Store struct {
	rds Client
}

// Client is a redis client for valkey.
//
// This is a minimal interface to allow for different redis clients.
type Client interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd
	io.Closer
}

// NewStore creates a new Store from a redis client.
func NewStore(rds Client) *Store {
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
	var (
		res  *redis.StringCmd
		val  string
		utPr semanticrouter.Utterance
	)
	res = s.rds.Get(ctx, utterance)
	val, err = res.Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("key does not exist: %w", err)
		}
		return nil, err
	}
	err = json.Unmarshal([]byte(val), &utPr)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling embedding: %w", err)
	}
	return utPr.Embed, nil
}

// Set sets a value in the valkey store.
func (s *Store) Set(
	ctx context.Context,
	utterance semanticrouter.Utterance,
) (err error) {
	var (
		val []byte
		res *redis.StatusCmd
	)
	val, err = json.Marshal(utterance)
	if err != nil {
		return fmt.Errorf("error marshaling embedding: %w", err)
	}
	res = s.rds.Set(
		ctx,
		utterance.Utterance,
		string(val),
		0,
	)
	err = res.Err()
	if err != nil {
		return fmt.Errorf("error setting embedding: %w", err)
	}
	return nil
}
