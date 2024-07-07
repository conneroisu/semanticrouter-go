package ollama

import (
	"context"
	"log"

	"github.com/ollama/ollama/api"
)

// Encoder is an encoder using Ollama models.
type Encoder struct {
	Client *api.Client
	Model  string
}

// NewEncoder creates a new Encoder.
func NewEncoder(client *api.Client, model string) *Encoder {
	return &Encoder{Client: client, Model: model}
}

// Encode encodes a query string into a Ollama embedding.
func (e *Encoder) Encode(ctx context.Context, query string) (result []float64, err error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		req := &api.EmbeddingRequest{
			Model:  e.Model,
			Prompt: query,
		}
		em, err := e.Client.Embeddings(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		result = em.Embedding
		return result, nil
	}
}
