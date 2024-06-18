package stores

import (
	"context"
	"fmt"
	"log"

	_ "github.com/asg017/sqlite-vss/bindings/go"
	"github.com/bytedance/sonic"
	"github.com/conneroisu/go-semantic-router/domain"
	"github.com/uptrace/bun"
)

// SqliteStore stores a value in the database
type SqliteStore struct {
	Database     *bun.DB
	SQLSetupFunc func(db *bun.DB) error
}

// Store stores a value in the database
func (s *SqliteStore) Store(ctx context.Context, key string, value []float64) error {
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
func (s *SqliteStore) Get(ctx context.Context, key string) ([]float64, error) {
	var embedding []float64
	var utterance domain.Utterance
	var err error
	_, err = s.Database.NewSelect().Model(&utterance).
		Where("utterance = ?", key).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting value: %w", err)
	}
	if utterance.Embedding == nil {
		var em interface{}
		_, err = s.Database.NewSelect().Model(&em).
			Where("utterance = ?", key).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("error getting value: %w", err)
		}
		embedding = em.([]float64)
	}
	return embedding, nil
}
