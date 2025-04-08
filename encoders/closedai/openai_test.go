package closedai

import (
	"context"
	"testing"

	openai "github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
)

func TestEncoder_Encode(t *testing.T) {
	t.Parallel()
	client := &mockClient{}
	encoder := NewEncoder(client, "text-embedding-ada-002")
	utterance := "Hello, world!"
	got, err := encoder.Encode(t.Context(), utterance)
	expected := []float64{
		0.0,
	}
	assert.Equal(t, expected, got)
	assert.NoError(t, err)
}

type mockClient struct{}

func (m *mockClient) CreateEmbeddings(
	_ context.Context,
	req openai.EmbeddingRequest,
) (openai.EmbeddingResponse, error) {
	return openai.EmbeddingResponse{
		Data: []openai.Embedding{
			{
				Embedding: []float32{
					0.0,
				},
			},
		},
	}, nil
}
