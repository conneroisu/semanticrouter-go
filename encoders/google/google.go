package google

import (
	"context"

	"github.com/google/generative-ai-go/genai"
)

// Client is a minimal client for the Google Generative AI API.
type Client interface {
	EmbeddingModel(name string) Model
}

// Model is a minimal model for the Google Generative AI API.
type Model interface {
	EmbedContent(ctx context.Context, content genai.Text) (genai.EmbedContentResponse, error)
}

// Encoder encodes a query string into a Google search URL.
type Encoder struct {
	client Client
	name   string
}

// NewEncoder creates a new GoogleEncoder.
func NewEncoder(
	client Client,
) *Encoder {
	return &Encoder{
		client: client,
	}
}

// Encode encodes a query string into a Google search URL.
func (e *Encoder) Encode(
	ctx context.Context,
	query string,
) ([]float64, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		embedding, err := e.client.EmbeddingModel(
			e.name,
		).EmbedContent(ctx, genai.Text(query))
		if err != nil {
			return nil, err
		}
		// type float32
		a := embedding.Embedding.Values
		// convert to []float64
		b := make([]float64, len(a))
		for i, v := range a {
			b[i] = float64(v)
		}
		return b, nil
	}
}
