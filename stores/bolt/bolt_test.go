package bolt_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	semanticrouter "github.com/conneroisu/semanticrouter-go"
	boltstore "github.com/conneroisu/semanticrouter-go/stores/bolt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoltStore(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "boltdb-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)
	
	// Create a temporary file for the test
	tempFile := filepath.Join(tempDir, "test-embeddings.db")

	// Create a new BoltDB store
	store, err := boltstore.NewStore(tempFile)
	require.NoError(t, err)
	defer store.Close()

	// Create a context for the operations
	ctx := context.Background()

	// Test storing an embedding
	testEmbedding := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	utterance := semanticrouter.Utterance{
		Utterance: "test utterance",
		Embed:     testEmbedding,
	}

	err = store.Set(ctx, utterance)
	assert.NoError(t, err)

	// Test retrieving the embedding
	retrieved, err := store.Get(ctx, utterance.Utterance)
	assert.NoError(t, err)
	assert.Equal(t, testEmbedding, retrieved)

	// Test retrieving a non-existent key
	_, err = store.Get(ctx, "non-existent")
	assert.Error(t, err)

	// Test custom bucket name
	customBucketFile := filepath.Join(tempDir, "custom-bucket.db")
	customBucketStore, err := boltstore.NewStore(
		customBucketFile,
		boltstore.WithBucketName("custom-bucket"),
	)
	require.NoError(t, err)
	defer customBucketStore.Close()

	err = customBucketStore.Set(ctx, utterance)
	assert.NoError(t, err)

	retrieved, err = customBucketStore.Get(ctx, utterance.Utterance)
	assert.NoError(t, err)
	assert.Equal(t, testEmbedding, retrieved)
}

func TestBoltStoreCancellation(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "boltdb-cancel-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)
	
	// Create a temporary file for the test
	tempFile := filepath.Join(tempDir, "test-embeddings-cancel.db")

	// Create a new BoltDB store
	store, err := boltstore.NewStore(tempFile)
	require.NoError(t, err)
	defer store.Close()

	// Create a cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Test storing with cancelled context
	testEmbedding := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	utterance := semanticrouter.Utterance{
		Utterance: "test utterance",
		Embed:     testEmbedding,
	}

	err = store.Set(ctx, utterance)
	assert.Error(t, err)

	// Test retrieving with cancelled context
	_, err = store.Get(ctx, utterance.Utterance)
	assert.Error(t, err)
}