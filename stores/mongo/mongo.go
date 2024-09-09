// Package mongo provides a MongoDB store for embeddings.
package mongo

import (
	"context"

	"github.com/conneroisu/semanticrouter-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Store is a MongoDB store.
//
// It implements the Store interface.
type Store struct {
	coll *mongo.Collection
}

// New creates a new MongoDB store.
func New(collection *mongo.Collection) *Store {
	return &Store{
		coll: collection,
	}
}

// Get gets a value from the store.
func (s *Store) Get(ctx context.Context, utterance string) ([]float64, error) {
	var floats []float64
	filter := bson.M{"utterance": utterance}
	cur, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []semanticrouter.Utterance
	if err = cur.All(ctx, &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		floats = append(floats, result.Embed...)
	}
	defer cur.Close(ctx)
	return floats, nil
}

// Set stores a value in the store.
func (s *Store) Set(ctx context.Context, keyValPair semanticrouter.Utterance) error {
	_, err := s.coll.InsertOne(ctx, keyValPair)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the MongoDB connection.
func (s *Store) Close() error {
	return nil
}
