package stores

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	minioContainer "github.com/testcontainers/testcontainers-go/modules/minio"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// TestStore is a test for the minio store.
func TestStore(t *testing.T) {
	ctx := context.Background()
	minioContainer, err := minioContainer.RunContainer(ctx, testcontainers.WithImage("minio/minio:RELEASE.2024-01-16T16-07-38Z"))

	assert.NoError(t, err)
	endpoint, err := minioContainer.Endpoint(ctx, "")
	assert.NoError(t, err)
	client, err := minio.New(
		endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4("Q3AM3UQ867SPQQA43P2F", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", ""),
			Secure: true,
		},
	)
	assert.NoError(t, err)
	store := Store{
		Client: client,
		Bucket: "testbucket",
	}
	err = store.Store(
		ctx,
		"key",
		[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
	)
	assert.NoError(t, err)

	floats, err := store.Get(ctx, "key")
	assert.NoError(t, err)
	assert.NotNil(t, floats)

	assert.Equal(
		t,
		[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
		floats,
	)
}
