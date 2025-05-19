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

// Test to ensure setting the same key overwrites the previous value.
func TestBoltStoreOverwriteExistingKey(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "overwrite-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	dbFile := filepath.Join(tempDir, "overwrite.db")
	store, err := boltstore.NewStore(dbFile)
	require.NoError(t, err)
	defer store.Close()

	ctx := context.Background()
	key := "overwrite-key"
	firstEmbed := []float64{1.1, 2.2}
	require.NoError(t, store.Set(ctx, semanticrouter.Utterance{Utterance: key, Embed: firstEmbed}))

	secondEmbed := []float64{3.3, 4.4}
	require.NoError(t, store.Set(ctx, semanticrouter.Utterance{Utterance: key, Embed: secondEmbed}))

	retrieved, err := store.Get(ctx, key)
	require.NoError(t, err)
	assert.Equal(t, secondEmbed, retrieved)
}

// Test storing and retrieving multiple distinct keys.
func TestBoltStoreMultipleKeys(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "multiple-keys-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	dbFile := filepath.Join(tempDir, "multiple.db")
	store, err := boltstore.NewStore(dbFile)
	require.NoError(t, err)
	defer store.Close()

	ctx := context.Background()
	entries := []semanticrouter.Utterance{
		{Utterance: "one", Embed: []float64{1.0}},
		{Utterance: "two", Embed: []float64{2.0, 2.1}},
		{Utterance: "three", Embed: []float64{3.0, 3.1, 3.2}},
	}
	for _, e := range entries {
		require.NoError(t, store.Set(ctx, e))
	}
	for _, e := range entries {
		retrieved, err := store.Get(ctx, e.Utterance)
		require.NoError(t, err)
		assert.Equal(t, e.Embed, retrieved)
	}
}

// Test setting and retrieving an empty embedding slice.
func TestBoltStoreEmptyEmbedding(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "empty-embed-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	dbFile := filepath.Join(tempDir, "empty.db")
	store, err := boltstore.NewStore(dbFile)
	require.NoError(t, err)
	defer store.Close()

	ctx := context.Background()
	entry := semanticrouter.Utterance{Utterance: "empty", Embed: []float64{}}
	require.NoError(t, store.Set(ctx, entry))

	retrieved, err := store.Get(ctx, entry.Utterance)
	require.NoError(t, err)
	assert.Empty(t, retrieved)
}

// Test behavior after closing the store: all operations should return an error.
func TestBoltStoreCloseBehavior(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "close-behavior-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	dbFile := filepath.Join(tempDir, "close.db")
	store, err := boltstore.NewStore(dbFile)
	require.NoError(t, err)
	require.NoError(t, store.Close())

	ctx := context.Background()
	entry := semanticrouter.Utterance{Utterance: "post-close", Embed: []float64{0.0}}

	err = store.Set(ctx, entry)
	assert.Error(t, err)

	_, err = store.Get(ctx, entry.Utterance)
	assert.Error(t, err)
}

// Test NewStore returns an error when the parent directory does not exist.
func TestBoltStoreNewStoreInvalidPath(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "invalid-path-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	invalidDir := filepath.Join(tempDir, "no-such-dir")
	invalidFile := filepath.Join(invalidDir, "db.db")
	_, err = boltstore.NewStore(invalidFile)
	assert.Error(t, err)
}