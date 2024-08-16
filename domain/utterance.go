package domain

import (
	"encoding/json"
	"fmt"
)

// Embedding is the embedding of some text, speech, or other data (images, videos, etc.).
type Embedding []float64

// Utterance represents a utterance in the semantic router.
type Utterance struct {
	// ID is the ID of the utterance.
	ID int
	// Utterance is the utterance.
	Utterance string
	// EmbeddingBytes is the embedding of the utterance.
	EmbeddingBytes []byte
	// Embed is the Embed of the utterance.
	Embed Embedding
}

// UtterancePrime represents a utterance in the semantic router.
type UtterancePrime struct {
	Utterance string    `json:"utterance,omitempty"`
	Embedding []float64 `json:"embedding"` // Embedding is the embedding of the utterance.
}

// SetEmbedding sets the embedding of the utterance.
func (u *Utterance) SetEmbedding(embedding []float64) error {
	type E struct {
		Embedding []float64 `json:"embedding"`
	}
	var embeddingBytes []byte
	embeddingBytes, err := json.Marshal(E{Embedding: embedding})
	if err != nil {
		return fmt.Errorf("error marshaling embedding: %w", err)
	}
	u.EmbeddingBytes = embeddingBytes
	return nil
}

// Embedding returns the embedding of the utterance.
func (u *Utterance) Embedding() (Embedding, error) {
	type E struct {
		Embedding []float64 `json:"embedding"`
	}
	var embedding E
	err := json.Unmarshal(u.EmbeddingBytes, &embedding)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling embedding: %w", err)
	}
	return embedding.Embedding, nil
}
