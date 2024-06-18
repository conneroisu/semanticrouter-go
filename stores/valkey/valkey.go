package valkey

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Store is a simple key-value store for embeddings.
type Store struct {
	rds *redis.Client
}

// Get gets a value from the
func (s *Store) Get(ctx context.Context, utterance string) ([]float64, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := s.rds.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

// Set sets a value in the store
func (s *Store) Set(ctx context.Context, utterance string, value []float64) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}
