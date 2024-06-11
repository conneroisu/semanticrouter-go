package semantic_router

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Routes represents a route in the semantic router.
type Routes struct {
	Name       string   `json:"name" yaml:"name" toml:"name"`
	Utterances []string `json:"utterances" yaml:"utterances" toml:"utterances"`
}

// Encoder represents a encoding driver in the semantic router.
type Encoder interface {
	Encode(string) ([]float64, error)
}

// Router represents a semantic router.
type Router struct {
	Routes  []Routes
	Encoder Encoder
}

// NewRouter creates a new semantic router.
func NewRouter(routes []Routes, encoder Encoder) *Router {
	return &Router{
		Routes:  routes,
		Encoder: encoder,
	}
}

// Match returns the route that matches the given utterance.
func (r *Router) Match(utterance string) (string, error) {
	encoding, err := r.Encoder.Encode(utterance)
	if err != nil {
		return "", fmt.Errorf("error encoding utterance: %w", err)
	}
	vecs := make([]float64, len(r.Routes))
	for _, route := range r.Routes {
		for _, utterance := range route.Utterances {
			vec, err := r.Encoder.Encode(utterance)
			if err != nil {
				return "", fmt.Errorf("error encoding utterance: %w", err)
			}
			for i, v := range vec {
				vecs[i] += v
			}
		}
	}
	sim := SimilarityMatrix(
		mat.NewVecDense(len(encoding), encoding),
		mat.NewDense(len(vecs), len(encoding), vecs),
	)
	scores, indices := TopScores(sim, 1)
	if len(scores) == 0 {
		return "", fmt.Errorf("no route found")
	}
	utterance = r.Routes[indices[0]].Utterances[0]
	return utterance, nil
}
