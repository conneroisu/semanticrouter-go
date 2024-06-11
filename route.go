package semantic_router

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

// Router represents a semantic router.
//
// Router is a struct that contains a slice of Routes and an Encoder.
//
// Match can be called on a Router to find the best route for a given utterance.
type Router struct {
	// Routes is a slice of Routes.
	Routes []Route
	// Encoder is an Encoder that encodes utterances into vectors.
	Encoder Encoder
}

// Route represents a route in the semantic router.
//
// It is a struct that contains a name and a slice of Utterances.
type Route struct {
	Name       string      `json:"name"       yaml:"name"       toml:"name"`
	Utterances []Utterance `json:"utterances" yaml:"utterances" toml:"utterances"`
}

// Utterance represents a utterance in the semantic router.
type Utterance struct {
	Utterance string    `json:"utterance" yaml:"utterance" toml:"utterance"`
	Embedding []float64 `json:"embedding" yaml:"embedding" toml:"embedding"`
}

// Encoder represents a encoding driver in the semantic router.
//
// It is an interface that defines a single method, Encode, which takes a string
// and returns a []float64 representing the embedding of the string.
type Encoder interface {
	Encode(string) ([]float64, error)
}

// NewRouter creates a new semantic router.
func NewRouter(routes []Route, encoder Encoder) (*Router, error) {
	routesLen := len(routes)
	for i := 0; i < routesLen; i++ {
		route := routes[i]
		utters := route.Utterances
		utterLen := len(utters)
		for j := 0; j < utterLen; j++ {
			utter := utters[j]
			colVec, err := encoder.Encode(utter.Utterance)
			if err != nil {
				return nil,
					fmt.Errorf(
						"error encoding utterance: %s: %w",
						utter.Utterance,
						err,
					)
			}
			route.Utterances[j].Embedding = colVec
		}
	}
	return &Router{
		Routes:  routes,
		Encoder: encoder,
	}, nil
}

// Match returns the route that matches the given utterance.
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
			if len(ut.Embedding) != queryVec.Len() {
				log.Printf(
					"Embedding length mismatch: queryVec.Len() = %d, ut.Embedding.Len() = %d",
					queryVec.Len(),
					len(ut.Embedding),
				)
				continue
			}
			indexVec := mat.NewVecDense(len(ut.Embedding), ut.Embedding)
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
