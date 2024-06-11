package providers

import (
	"context"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

// OpenAIEncoder encodes a query string into an OpenAI embedding.
type OpenAIEncoder struct {
	Ctx    context.Context
	APIKey string
	Model  openai.EmbeddingModel
}

// Encode encodes the given utterance using the OpenAI API.
func (o OpenAIEncoder) Encode(utterance string) ([]float64, error) {
	client := openai.NewClient(o.APIKey)
	queryReq := openai.EmbeddingRequest{
		Input: utterance,
		Model: openai.AdaEmbeddingV2,
	}
	queryResponse, err := client.CreateEmbeddings(
		context.Background(),
		queryReq,
	)
	if err != nil {
		log.Fatal("Error creating query embedding:", err)
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
