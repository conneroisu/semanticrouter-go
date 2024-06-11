package providers

import (
	"context"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIEncoder struct {
	APIKey string
	Model  openai.EmbeddingModel
}

// Encode encodes the given utterance using the OpenAI API.
func (o *OpenAIEncoder) Encode(utterance string) ([]float32, error) {
	client := openai.NewClient(o.APIKey)
	queryReq := openai.EmbeddingRequest{
		Input: utterance,
		Model: openai.AdaEmbeddingV2,
	}
	// Create an embedding for the user query
	queryResponse, err := client.CreateEmbeddings(context.Background(), queryReq)
	if err != nil {
		log.Fatal("Error creating query embedding:", err)
	}
	var floats []float32
	for _, f := range queryResponse.Data[0].Embedding {
		floats = append(floats, float32(f))
	}
	return floats, nil
}
