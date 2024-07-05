package ollama

import (
	"context"
	"log"
	"testing"

	"github.com/ollama/ollama/api"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	tcollama "github.com/testcontainers/testcontainers-go/modules/ollama"
)

// TestEncoder_Encode tests the Encode method of the Encoder struct.
func TestEncoder_Encode(t *testing.T) {
	ctx := context.Background()
	ollamaContainer, err := tcollama.RunContainer(ctx, testcontainers.WithImage("ollama/ollama:0.1.25"))
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}
	// Clean up the container
	defer func() {
		if err := ollamaContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err) // nolint:gocritic
		}
	}()
	cli, err := api.ClientFromEnvironment()
	assert.NoError(t, err)
	encoder := Encoder{
		Client: cli,
		Model:  "mxbai-embed-large",
	}
	result, err := encoder.Encode("hello world")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
