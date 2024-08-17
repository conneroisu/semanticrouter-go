package mongo_test

import (
	"context"
	"log"
	"testing"

	semanticrouter "github.com/conneroisu/go-semantic-router"
	"github.com/conneroisu/go-semantic-router/stores/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

var (
	_ semanticrouter.Store = (*mongo.Store)(nil)
)

func TestStore(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	mongodbContainer, err := mongodb.Run(ctx, "mongo:6")
	a.NoError(err)
	defer func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()
	uri, err := mongodbContainer.ConnectionString(ctx)
	a.NoError(err)
	store, err := mongo.New(uri, "test", "test")
	a.NoError(err)
	err = store.Store(
		ctx,
		semanticrouter.Utterance{
			Utterance: "key",
			Embed:     []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		},
	)
	a.NoError(err)

	floats, err := store.Get(ctx, "key")
	a.NoError(err)
	a.Len(floats, 5)
}
