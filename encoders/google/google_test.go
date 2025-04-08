package google

import (
	"context"
	"testing"

	"github.com/google/generative-ai-go/genai"
	"github.com/stretchr/testify/assert"
)

func TestEncoder_Encode(t *testing.T) {
	ctx := context.Background()
	mockClient := mockClient{}
	encoder := NewEncoder(mockClient)
	result, err := encoder.Encode(ctx, "query")
	assert.NoError(t, err)
	assert.Equal(t, []float64{0.0}, result)
}

type mockClient struct{}

func (m mockClient) EmbeddingModel(_ string) Model {
	return mockModel{}
}

type mockModel struct{}

func (m mockModel) EmbedContent(_ context.Context, _ genai.Text) (genai.EmbedContentResponse, error) {
	return genai.EmbedContentResponse{
		Embedding: &genai.ContentEmbedding{
			Values: []float32{0.0},
		},
	}, nil
}
