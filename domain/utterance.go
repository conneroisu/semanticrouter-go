package domain

import (
	"github.com/bytedance/sonic"
	"github.com/uptrace/bun"
)

type Embedding []float64

// Utterance represents a utterance in the semantic router.
type Utterance struct {
	*bun.BaseModel `bun:"table:utterances"`
	// ID is the ID of the utterance.
	ID int `bun:"id,pk,autoincrement"`
	// Utterance is the utterance.
	Utterance string `bun:"utterance"` // Utterance is the utterance.
	// EmbeddingBytes is the embedding of the utterance.
	EmbeddingBytes []byte `bun:"embedding" json:"embedding"` // Embedding is the embedding of the utterance.
	// Embed is the Embed of the utterance.
	Embed Embedding
}

// UtterancePrime represents a utterance in the semantic router.
type UtterancePrime struct {
	Embedding []float64 `json:"embedding"` // Embedding is the embedding of the utterance.
}

// Embedding returns the embedding of the utterance.
func (u *Utterance) Embedding() (Embedding, error) {
	type Embedding struct {
		Embedding []float64 `json:"embedding"`
	}
	var embedding Embedding
	err := sonic.Unmarshal(u.EmbeddingBytes, &embedding)
	if err != nil {
		return nil, err
	}
	return embedding.Embedding, nil
}
