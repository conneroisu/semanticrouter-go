package ollama_test

import (
	"context"
	"testing"

	"github.com/conneroisu/go-semantic-router/encoders/ollama"
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
	encoder := ollama.NewEncoder(client, "all-minilm")
	result, err := encoder.Encode(ctx, "hello world")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
