package voyageai

import (
	"fmt"

	"github.com/conneroisu/go-voyageai"
)

// Encoder is an encoder using VoyageAI embedding models.
type Encoder struct {
	Client *voyageai.Client
	Model  string
}

// NewEncoder creates a new Encoder.
func NewEncoder(client *voyageai.Client, model string) *Encoder {
	return &Encoder{Client: client, Model: model}
}

// Encode encodes a query string into a VoyageAI embedding.
func (e *Encoder) Encode(query string) (result []float64, err error) {
	resp, err := e.Client.Embeddings(voyageai.EmbeddingsRequest{
		Model: e.Model,
		Input: []string{query},
	})
	if err != nil {
		return nil, fmt.Errorf("error creating query embedding: %w", err)
	}
	return resp.Data[0].Embedding, nil
}
