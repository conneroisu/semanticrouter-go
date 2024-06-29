package semanticrouter

import (
	"context"
	"fmt"

	"github.com/conneroisu/go-semantic-router/domain"
	"gonum.org/v1/gonum/mat"
)

// Router represents a semantic router.
//
// Router is a struct that contains a slice of Routes and an Encoder.
//
// Match can be called on a Router to find the best route for a given utterance.
type Router struct {
	Routes  []Route `json:"routes" yaml:"routes" toml:"routes"`    // Routes is a slice of Routes.
	Encoder Encoder `json:"encoder" yaml:"encoder" toml:"encoder"` // Encoder is an Encoder that encodes utterances into vectors.
	Storage Store   `json:"storage" yaml:"storage" toml:"storage"` // Storage is a Store that stores the utterances.
}

// Route represents a route in the semantic router.
//
// It is a struct that contains a name and a slice of Utterances.
type Route struct {
	Name       string             `json:"name"       yaml:"name"       toml:"name"`       // Name is the name of the route.
	Utterances []domain.Utterance `json:"utterances" yaml:"utterances" toml:"utterances"` // Utterances is a slice of Utterances.
}

// Encoder represents a encoding driver in the semantic router.
//
// It is an interface that defines a single method, Encode, which takes a string
// and returns a []float64 representing the embedding of the string.
type Encoder interface {
	Encode(string) ([]float64, error)
}

// Store is an interface that defines a method, Store, which takes a []float64
// and stores it in a some sort of data store, and a method, Get, which takes a
// string and returns a []float64 from the data store.
type Store interface {
	Store(ctx context.Context, utterance domain.Utterance) error
	Get(ctx context.Context, utterance string) ([]float64, error)
}

// NewRouter creates a new semantic router.
func NewRouter(routes []Route, encoder Encoder, store Store) (*Router, error) {
	routesLen := len(routes)
	ctx := context.Background()
	for i := 0; i < routesLen; i++ {
		route := routes[i]
		utters := route.Utterances
		for _, utter := range utters {
			embedding, err := store.Get(ctx, utter.Utterance)
			if err == nil && len(embedding) != 0 {
				continue
			}
			err = store.Store(ctx, utter)
			if err != nil {
				return nil,
					fmt.Errorf(
						"error storing utterance: %s: %w",
						utter.Utterance,
						err,
					)
			}
		}
	}
	return &Router{
		Routes:  routes,
		Encoder: encoder,
		Storage: store,
	}, nil
}

// Match returns the route that matches the given utterance.
// The score is the similarity score between the query vector and the index vector.
func (r *Router) Match(utterance string) (string, float64, error) {
	encoding, err := r.Encoder.Encode(utterance)
	if err != nil {
		return "", 0.0, fmt.Errorf("error encoding utterance: %w", err)
	}
	bestScore := 0.0
	var bestRouteName string
	queryVec := mat.NewVecDense(len(encoding), encoding)
	for _, route := range r.Routes {
		for _, ut := range route.Utterances {
			em, err := ut.Embedding()
			if err != nil {
				return "", 0.0, fmt.Errorf("error getting embedding: %w", err)
			}
			emLen := len(em)
			if emLen != queryVec.Len() {
				continue
			}
			indexVec := mat.NewVecDense(emLen, em)
			simScore := SimilarityMatrix(queryVec, indexVec)
			if simScore > bestScore {
				bestScore = simScore
				bestRouteName = route.Name
			}
		}
	}
	if bestRouteName == "" {
		return "", 0.0, fmt.Errorf("no route found")
	}
	return bestRouteName, bestScore, nil
}
