package semanticrouter

import (
	"context"
	"io"

	"gonum.org/v1/gonum/mat"
)

// Encoder represents a encoding driver in the semantic router.
//
// It is an interface that defines a single method, Encode, which takes a string
// and returns a []float64 representing the embedding of the string.
type Encoder interface {
	Encode(ctx context.Context, utterance string) ([]float64, error)
}

// Store is an interface that defines a method, Store, which takes a []float64
// and stores it in a some sort of data store, and a method, Get, which takes a
// string and returns a []float64 from the data store.
type Store interface {
	Setter
	Getter
	io.Closer
}

// Setter is an interface that defines a method, Store, which takes a []float64
// and stores it in a some sort of data store.
type Setter interface {
	Set(ctx context.Context, keyValPair Utterance) error
}

// Getter is an interface that defines a method, Get, which takes a
// string and returns a []float64 from the data store.
//
// If the key does not exist, it returns an error.
type Getter interface {
	Get(ctx context.Context, key string) ([]float64, error)
}

// Option is a function that configures a Router.
type Option func(*Router)

// handler is a function that takes two vectors and returns a float64.
//
// It also returns an error if there is an error during the comparison.
//
// It is used to compare the similarity between two vectors.
type handler func(queryVec *mat.VecDense, indexVec *mat.VecDense) (float64, error)
