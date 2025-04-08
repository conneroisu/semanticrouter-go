package mongo

import (
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
	if testing.Short() {
		t.Skip("skipping non-short test")
	}
	a := assert.New(t)
	mongodbContainer, err := mongodb.Run(t.Context(), "mongo:6")
	a.NoError(err)
	defer func() {
		if Terr := mongodbContainer.Terminate(t.Context()); Terr != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()
	uri, err := mongodbContainer.ConnectionString(t.Context())
	a.NoError(err)
	client, err := mongo.Connect(t.Context(), options.Client().ApplyURI(uri))
	a.NoError(err)
	defer func() {
		err = client.Disconnect(t.Context())
		if err != nil {
			log.Fatalf("failed to disconnect from mongodb: %s", err)
		}
	}()
	collection := client.Database("test").Collection("test")
	store := New(collection)
	a.NoError(err)
	err = store.Set(
		t.Context(),
		semanticrouter.Utterance{
			Utterance: "key",
			Embed:     []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		})
	a.NoError(err)

	floats, err := store.Get(t.Context(), "key")
	a.NoError(err)
	a.Len(floats, 5)
}
