// Package ollama provides an encoder using Ollama models.
//
// Ollama is the easiest way to get started with LLMs.
package ollama

import (
	"context"
	"log"

	"github.com/ollama/ollama/api"
)

// Encoder is an encoder using Ollama models.
type Encoder struct {
	Client *api.Client
}

// NewEncoder creates a new Encoder.
func NewEncoder(client *api.Client) *Encoder {
	return &Encoder{Client: client}
}

// Encode encodes a query string into a Ollama embedding.
func (e *Encoder) Encode(query string) (result []float64, err error) {
	req := &api.EmbeddingRequest{
		Model:  "mxbai-embed-large",
		Prompt: query,
	}
	ctx := context.Background()
	em, err := e.Client.Embeddings(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	result = em.Embedding
	return result, nil
}
