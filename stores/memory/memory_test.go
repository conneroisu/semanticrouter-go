package memory_test

import (
	"context"
	"testing"

	"github.com/conneroisu/semanticrouter-go"
	"github.com/conneroisu/semanticrouter-go/stores/memory"
	"github.com/stretchr/testify/assert"
)

var (
	_ semanticrouter.Store = (*memory.Store)(nil)
)

// TestStore tests the in memory store.
func TestStore(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	store := memory.NewStore()
	utter := semanticrouter.Utterance{
		Utterance: "key",
	}
	utter.Embed = []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	err := store.Set(
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
