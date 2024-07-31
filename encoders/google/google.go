package encoders

import (
	"context"

	"github.com/google/generative-ai-go/genai"
)

// GoogleEncoder encodes a query string into a Google search URL.
type GoogleEncoder struct {
	client genai.Client
	name   string
}

// NewGoogleEncoder creates a new GoogleEncoder.
func NewGoogleEncoder(
	client genai.Client,
) *GoogleEncoder {
	return &GoogleEncoder{
		client: client,
	}
}

// Encode encodes a query string into a Google search URL.
func (e *GoogleEncoder) Encode(
	ctx context.Context,
	query string,
) ([]float64, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		model := e.client.EmbeddingModel(e.name)
		embedding, err := model.EmbedContent(ctx, genai.Text(query))
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
