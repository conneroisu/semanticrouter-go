package mongo_test

import (
	"context"
	"log"
	"testing"

	"github.com/conneroisu/go-semantic-router/stores/mongo"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func TestStore(t *testing.T) {
	ctx := context.Background()

	mongodbContainer, err := mongodb.Run(ctx, "mongo:6")
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	// Clean up the container
	defer func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	// Get the MongoDB URI
	uri, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("failed to get connection string: %s", err)
	}

	// Create a new MongoDB store
	store, err := mongo.New(uri, "test", "test")
	if err != nil {
		log.Fatalf("failed to create store: %s", err)
	}

	err = store.Store(
		ctx,
		"key",
		[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
	)
	if err != nil {
		log.Fatalf("failed to set key: %s", err)
	}

	floats, err := store.Get(ctx, "key")
	if err != nil {
		log.Fatalf("failed to get key: %s", err)
	}
	if len(floats) != 5 {
		log.Fatalf("expected length of floats to be 5, got %d", len(floats))
	}
}
