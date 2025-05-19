package bolt

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/conneroisu/semanticrouter-go"
	bolt "go.etcd.io/bbolt"
)

// Store is a BoltDB implementation of the semanticrouter Store interface.
type Store struct {
	db         *bolt.DB
	bucketName []byte
}

// StoreOption is a function that configures a Store.
type StoreOption func(*Store)

// WithBucketName sets the bucket name for the store.
func WithBucketName(name string) StoreOption {
	return func(s *Store) {
		s.bucketName = []byte(name)
	}
}

// NewStore creates a new BoltDB store for embeddings.
// If the database file doesn't exist, it will be created.
func NewStore(path string, opts ...StoreOption) (*Store, error) {
	// Ensure the directory exists
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create directory for BoltDB: %w", err)
		}
	}

	// Open the BoltDB database
	db, err := bolt.Open(path, 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open BoltDB: %w", err)
	}

	// Create store with default bucket name
	store := &Store{
		db:         db,
		bucketName: []byte("embeddings"),
	}

	// Apply options
	for _, opt := range opts {
		opt(store)
	}

	// Create the bucket if it doesn't exist
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(store.bucketName)
		if err != nil {
			return fmt.Errorf("failed to create bucket: %w", err)
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}

	return store, nil
}

// Set stores an embedding with the given key.
func (s *Store) Set(ctx context.Context, keyValPair semanticrouter.Utterance) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucketName)
		if b == nil {
			return fmt.Errorf("bucket %s not found", string(s.bucketName))
		}

		// Marshal embedding to JSON for storage
		data, err := json.Marshal(keyValPair.Embed)
		if err != nil {
			return fmt.Errorf("failed to marshal embedding: %w", err)
		}

		// Store the embedding with the utterance as the key
		return b.Put([]byte(keyValPair.Utterance), data)
	})
}

// Get retrieves an embedding for the given key.
func (s *Store) Get(ctx context.Context, key string) ([]float64, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var embedding []float64
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucketName)
		if b == nil {
			return fmt.Errorf("bucket %s not found", string(s.bucketName))
		}

		data := b.Get([]byte(key))
		if data == nil {
			return fmt.Errorf("key not found: %s", key)
		}

		return json.Unmarshal(data, &embedding)
	})

	if err != nil {
		return nil, err
	}

	return embedding, nil
}

// Close closes the underlying BoltDB database.
func (s *Store) Close() error {
	return s.db.Close()
}