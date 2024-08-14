package mongo

import (
	"context"
	"encoding/json"

	"github.com/conneroisu/go-semantic-router/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store is a MongoDB store.
type Store struct {
	mdb  *mongo.Client
	coll *mongo.Collection
}

// New creates a new MongoDB store.
func New(uri string, db, collection string) (*Store, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &Store{
		mdb:  client,
		coll: client.Database(db).Collection(collection),
	}, nil
}

// Get gets a value from the store.
func (s *Store) Get(ctx context.Context, utterance string) ([]float64, error) {
	var floats []float64
	filter := bson.M{"utterance": utterance}
	cur, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var doc bson.M
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		floats = append(floats, doc["embedding"].([]float64)...)
	}
	return floats, nil
}

// Store stores a value in the store.
func (s *Store) Store(ctx context.Context, utterance string, value []float64) error {
	jsonValue, err := json.Marshal(domain.UtterancePrime{Embedding: value})
	if err != nil {
		return err
	}
	filter := bson.M{"utterance": utterance}
	update := bson.M{"$set": bson.M{"utterance": utterance, "embedding": jsonValue}}

	result, err := s.coll.InsertOne(ctx,
		bson.D{
			{Key: "utterance", Value: utterance},
			{Key: "embedding", Value: jsonValue},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the MongoDB connection.
func (s *Store) Close() error {
	err := s.mdb.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}
