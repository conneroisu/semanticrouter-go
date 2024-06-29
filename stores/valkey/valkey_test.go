package valkey

import (
	"context"
	"fmt"
	"testing"

	clientLib "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// TestStore is a test for the redis/valkey store.
func TestStore(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "redis:7.2",
		ExposedPorts: []string{
			"6379/tcp",
		},
		WaitingFor: wait.ForLog("Ready to accept connections"),
	}
	redisContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	assert.NoError(t, err)
	endpoint, err := redisContainer.Endpoint(ctx, "")
	assert.NoError(t, err)
	store := Store{
		rds: clientLib.NewClient(
			&clientLib.Options{
				Addr:    endpoint,
				Network: "tcp",
			}),
	}

	val, err := store.Set(
		ctx,
		"key",
		[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
	)
	t.Log(fmt.Sprintf("val: %v", val))
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
