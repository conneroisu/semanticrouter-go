package sql

import (
	"context"
	"fmt"
	"log"
	"sync"

	_ "github.com/asg017/sqlite-vss/bindings/go"
	"github.com/bytedance/sonic"
	"github.com/conneroisu/go-semantic-router/domain"
	"github.com/uptrace/bun"
)

// Store stores a value in a sql database
//
// Dialects supported: sqlite, mysql, postgres, mssql
type Store struct {
	// WriteMutex is an optional mutex that can be used to lock the database for writing
	WriteMutex *sync.Mutex

	// Database is the database connection
	Database *bun.DB

	// SetupFunc is a function that is called when the database is initialized
	SetupFunc func(db *bun.DB) error
}

// Store stores a value in the database
func (s *Store) Store(ctx context.Context, key string, value []float64) error {
	if s.WriteMutex != nil {
		s.WriteMutex.Lock()
		defer s.WriteMutex.Unlock()
	}
	var err error

	var embedding []byte
	embedding, err = sonic.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.Database.ExecContext(ctx, "select vss_version()")
	if err != nil {
		return fmt.Errorf("error getting version: %w", err)
	}
	_, err = s.Database.NewCreateTable().Model(
		&domain.Utterance{}).IfNotExists().Exec(ctx)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	_, err = s.Database.NewInsert().Model(&domain.Utterance{
		Utterance:      key,
		EmbeddingBytes: embedding,
	}).Returning("*").Exec(ctx)
	if err != nil {
		return fmt.Errorf("error inserting value: %w", err)
	}
	return nil
}

// Get retrieves a value from the Database
func (s *Store) Get(ctx context.Context, key string) ([]float64, error) {
	var embedding []float64
	var utterance domain.Utterance
	var err error
	_, err = s.Database.NewSelect().Model(&utterance).
		Where("utterance = ?", key).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting value: %w", err)
	}
	if utterance.Embed == nil {
		var em domain.Embedding
		em, err = utterance.Embedding()
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling embedding: %w", err)
		}
		embedding = em
	}
	return embedding, nil
}
