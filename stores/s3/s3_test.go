package s3

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestAccessS3(t *testing.T) {
	awsEndpoint := "http://localhost:4566"
	awsRegion := "us-east-1"

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "localstack/localstack",
		ExposedPorts: []string{"4566:4566/tcp"},
		WaitingFor:   wait.ForLog("Ready"),
	}

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	defer func(c testcontainers.Container, ctx context.Context) {
		if err := c.Terminate(ctx); err != nil {
			fmt.Println(err)
		}
	}(c, ctx)

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(awsRegion), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		t.Error(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	bucketName := "test"
	if _, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: &bucketName,
	}); err != nil {
		t.Error(err)
	}

	list, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(list.Buckets), 1)
	assert.Equal(t, *list.Buckets[0].Name, bucketName)
}
