// Package mongo provides a MongoDB store for embeddings.
package mongo

import (
	"context"
	"io"

	"github.com/conneroisu/semanticrouter-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Cursor is a MongoDB cursor.
//
// It implements an minimal subset of the mongo.Cursor interface.
type Cursor interface {
	All(ctx context.Context, result any) error
	io.Closer
}

// Collection is a MongoDB collection.
//
// It implements an minimal subset of the mongo.Collection interface.
type Collection interface {
	Find(ctx context.Context, filter any, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
	InsertOne(ctx context.Context, doc any, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

// Store is a MongoDB store.
//
// It implements the Store interface.
type Store struct {
	coll Collection
}

// New creates a new MongoDB store.
func New(collection Collection) *Store {
	return &Store{coll: collection}
}

// Get gets a value from the store.
func (s *Store) Get(ctx context.Context, utterance string) ([]float64, error) {
	var (
		floats  []float64
		results []semanticrouter.Utterance
	)
	cur, err := s.coll.Find(ctx, bson.M{"utterance": utterance})
	if err != nil {
		return nil, err
	}
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
