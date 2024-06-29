package memory

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	ctx := context.Background()
	store := NewStore()
	val, err := store.Set(
		ctx,
		"key",
		[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
	)
	t.Log(fmt.Sprintf("val: %v", val))
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
