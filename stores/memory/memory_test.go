package memory

import (
	"context"
	"testing"

	"github.com/conneroisu/go-semantic-router/domain"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	ctx := context.Background()
	store := NewStore()
	utter := domain.Utterance{
		Utterance: "key",
	}
	utter.SetEmbedding([]float64{1.0, 2.0, 3.0, 4.0, 5.0})
	err := store.Store(
		ctx,
		utter,
	)
	assert.NoError(t, err)

	floats, err := store.Get(ctx, "key")
	assert.NoError(t, err)
	assert.NotNil(t, floats)

	assert.Equal(
		t,
		[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
		floats,
	)
}
