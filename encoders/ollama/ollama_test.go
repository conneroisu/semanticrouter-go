package ollama

import (
	"context"
	"testing"

	"github.com/ollama/ollama/api"
	"github.com/stretchr/testify/assert"
)

// TestEncoder tests the encoder.
func TestEncoder(t *testing.T) {
	ctx := context.Background()
	client, err := api.ClientFromEnvironment()
	if err != nil {
		t.Fatal(err)
	}
	encoder := NewEncoder(client, "all-minilm")
	result, err := encoder.Encode(ctx, "hello world")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
