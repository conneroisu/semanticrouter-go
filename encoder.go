package semanticrouter

import (
	"context"

	"github.com/conneroisu/go-semantic-router/domain"
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
	Store(ctx context.Context, keyValPair domain.Utterance) error
	Get(ctx context.Context, key string) ([]float64, error)
	Close() error
}

// Option is a function that configures a Router.
type Option func(*Router)

// biFuncCoefficient is an struct that represents a function and it's coefficient.
type biFuncCoefficient struct {
	Func        func(queryVec *mat.VecDense, indexVec *mat.VecDense) float64
	Coefficient float64
}

// WithSimilarityDotMatrix sets the similarity function to use with a coefficient.
func WithSimilarityDotMatrix(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoefficients = append(r.biFuncCoefficients, biFuncCoefficient{
			Func:        SimilarityDotMatrix,
			Coefficient: coefficient,
		})
	}
}

// WithEuclideanDistance sets the EuclideanDistance function with a coefficient.
func WithEuclideanDistance(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoefficients = append(r.biFuncCoefficients, biFuncCoefficient{
			Func:        EuclideanDistance,
			Coefficient: coefficient,
		})
	}
}

// WithManhattanDistance sets the ManhattanDistance function with a coefficient.
func WithManhattanDistance(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoefficients = append(r.biFuncCoefficients, biFuncCoefficient{
			Func:        ManhattanDistance,
			Coefficient: coefficient,
		})
	}
}

// WithJaccardSimilarity sets the JaccardSimilarity function with a coefficient.
func WithJaccardSimilarity(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoefficients = append(r.biFuncCoefficients, biFuncCoefficient{
			Func:        JaccardSimilarity,
			Coefficient: coefficient,
		})
	}
}

// WithPearsonCorrelation sets the PearsonCorrelation function with a coefficient.
func WithPearsonCorrelation(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoefficients = append(r.biFuncCoefficients, biFuncCoefficient{
			Func:        PearsonCorrelation,
			Coefficient: coefficient,
		})
	}
}
