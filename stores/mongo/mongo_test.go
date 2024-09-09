package mongo

import (
	"context"
	"log"
	"testing"

	"github.com/conneroisu/semanticrouter-go"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	_ semanticrouter.Store = (*Store)(nil)
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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	a.NoError(err)
	defer func() {
		err = client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("failed to disconnect from mongodb: %s", err)
		}
	}()
	collection := client.Database("test").Collection("test")
	store := New(client, collection)
	a.NoError(err)
	err = store.Set(
		ctx,
		semanticrouter.Utterance{
			Utterance: "key",
			Embed:     []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		})
	a.NoError(err)

	floats, err := store.Get(ctx, "key")
	a.NoError(err)
	a.Len(floats, 5)
}
