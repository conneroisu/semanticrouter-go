package memory

import (
	"context"
	"testing"

	"github.com/conneroisu/go-semantic-router/domain"
	"github.com/stretchr/testify/assert"
)

// TestStore tests the in memory store.
func TestStore(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	store := NewStore()
	utter := domain.Utterance{
		Utterance: "key",
	}
	err := utter.SetEmbedding([]float64{1.0, 2.0, 3.0, 4.0, 5.0})
	a.NoError(err)
	err = store.Store(
		ctx,
		utter,
	)
	a.NoError(err)

	floats, err := store.Get(ctx, "key")
	a.NoError(err)
	a.NotNil(floats)

	assert.Equal(
		t,
		[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
		floats,
	)
}
