package semanticrouter

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
	"gonum.org/v1/gonum/mat"
)

// Router represents a semantic router.
//
// Router is a struct that contains a slice of Routes and an Encoder.
//
// Match can be called on a Router to find the best route for a given utterance.
type Router struct {
	Routes      []Route             // Routes is a slice of Routes.
	Encoder     Encoder             // Encoder is an Encoder that encodes utterances into vectors.
	Storage     Store               // Storage is a Store that stores the utterances.
	biFuncCoeff []biFuncCoefficient // biFuncCoefficients is a slice of biFuncCoefficients that represent the bi-function coefficients.
	workers     int                 // workers is the number of workers to use for computing similarity scores.
}

// WithWorkers sets the number of workers to use for computing similarity scores.
func WithWorkers(workers int) Option {
	return func(r *Router) {
		r.workers = workers
	}
}

// Route represents a route in the semantic router.
//
// It is a struct that contains a name and a slice of Utterances.
type Route struct {
	Name       string      // Name is the name of the route.
	Utterances []Utterance // Utterances is a slice of Utterances.
	chatter    Chatter     // Chatter is the chatter to use for the route when MatchChat is called.
}

// biFuncCoefficient is an struct that represents a function and it's coefficient.
type biFuncCoefficient struct {
	handler     handler
	coefficient float64
}

// NewRouter creates a new semantic router.
func NewRouter(
	routes []Route,
	encoder Encoder,
	store Store,
	opts ...Option,
) (router *Router, err error) {
	router = &Router{}
	routesLen := len(routes)
	ctx := context.Background()
	if len(opts) == 0 {
		opts = []Option{
			WithSimilarityDotMatrix(1.0),
			WithEuclideanDistance(1.0),
			WithManhattanDistance(1.0),
			WithJaccardSimilarity(1.0),
			WithPearsonCorrelation(1.0),
			WithWorkers(1),
		}
	}
	for _, opt := range opts {
		opt(router)
	}
	for i := 0; i < routesLen; i++ {
		for _, utter := range routes[i].Utterances {
			_, err = store.Get(ctx, utter.Utterance)
			if err == nil {
				continue
			}
			en, err := encoder.Encode(ctx, utter.Utterance)
			if err != nil {
				return nil, fmt.Errorf("error encoding utterance: %w", err)
			}
			utter.Embed = en
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
//
// The score is the similarity score between the query vector and the index vector.
//
// If the given context is canceled, the context's error is returned if it is non-nil.
func (r *Router) Match(
	ctx context.Context,
	utterance string,
) (bestRoute *Route, bestScore float64, err error) {
	encoding, err := r.Encoder.Encode(ctx, utterance)
	if err != nil {
		return nil, 0.0, ErrEncoding{
			Message: fmt.Sprintf(
				"error encoding utterance: %s",
				utterance,
			),
		}
	}
	queryVec := mat.NewVecDense(len(encoding), encoding)
	for _, route := range r.Routes {
		for _, ut := range route.Utterances {
			em, err := r.Storage.Get(ctx, ut.Utterance)
			if err != nil {
				return nil, 0.0, ErrGetEmbedding{
					Message: fmt.Sprintf(
						"error getting embedding: %s",
						ut.Utterance,
					),
				}
			}
			emLen := len(em)
			if emLen != queryVec.Len() {
				continue
			}
			indexVec := mat.NewVecDense(emLen, em)
			simScore, err := r.computeScore(queryVec, indexVec)
			if err != nil {
				return nil, 0.0, err
			}
			if simScore > bestScore {
				bestScore = simScore
				bestRoute = &route
			}
		}
	}
	return bestRoute, bestScore, nil
}

// MatchChat returns the route that matches the given chat.
//
// The score is the similarity score between the query vector and the index vector.
//
// If the given context is canceled, the context's error is returned if it is non-nil.
func (r *Router) MatchChat(
	ctx context.Context,
	utterance string,
) (*string, error) {
	route, _, err := r.Match(ctx, utterance)
	if err != nil {
		return nil, err
	}
	if route.chatter == nil {
		return nil, fmt.Errorf("route %s has no chatter", route.Name)
	}
	resp, err := route.chatter.Chat(ctx, utterance)
	return &resp, err
}

// computeScore computes the score for a given utterance and route.
//
// It takes a query vector and an index vector as input and returns a score.
//
// Additionally, it leverages the router's biFuncCoefficients to apply different
// weighting factors to functions to get the similarity score.
func (r *Router) computeScore(
	queryVec *mat.VecDense,
	indexVec *mat.VecDense,
) (float64, error) {
	score := 0.0
	eg := errgroup.Group{}
	eg.SetLimit(r.workers)
	for _, fn := range r.biFuncCoeff {
		eg.Go(func() error {
			interScore, err := fn.handler(queryVec, indexVec)
			if err != nil {
				return err
			}
			score += fn.coefficient * interScore
			return nil
		})
	}
	return score, eg.Wait()
}
