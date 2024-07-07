package openai

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

// Encoder encodes a query string into an OpenAI embedding.
type Encoder struct {
	// Client is the OpenAI client.
	Client *openai.Client
	// Model is the OpenAI embedding model to use.
	Model openai.EmbeddingModel
}

// Encode encodes the given utterance using the OpenAI API.
func (o Encoder) Encode(ctx context.Context, utterance string) ([]float64, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		queryReq := openai.EmbeddingRequest{
			Input: utterance,
			Model: o.Model,
		}
		queryResponse, err := o.Client.CreateEmbeddings(
			context.Background(),
			queryReq,
		)
		if err != nil {
			return nil, fmt.Errorf("error creating query embedding: %w", err)
		}
		var floats []float32
		for _, f := range queryResponse.Data[0].Embedding {
			floats = append(floats, float32(f))
		}
		// Convert the float32 slice to a float64 slice
		floats64 := make([]float64, len(floats))
		for i, f := range floats {
			floats64[i] = float64(f)
		}
		return floats64, nil
	}
}
