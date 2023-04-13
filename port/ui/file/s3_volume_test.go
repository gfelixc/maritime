package file

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/require"
)

func TestS3Volume(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	region := "us-west-1"
	bucket := "test"
	endpoint := "http://localhost:4566"

	require.NoError(t, setUp(ctx, bucket, endpoint))

	volume, err := NewS3VolumeCustomEndpoint(ctx, region, bucket, endpoint)
	require.NoError(t, err)

	file, err := volume.Open(ctx, "file.txt")
	defer file.Close()
	require.NoError(t, err)

	data, err := io.ReadAll(file)
	require.NoError(t, err)
	require.Equal(t, "lorem ipsum", string(data))
}

func setUp(ctx context.Context, bucketName string, endpoint string) error {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithEndpointResolverWithOptions(CustomEndpointResolver{endpoint: endpoint}),
		config.WithCredentialsProvider(aws.AnonymousCredentials{}),
	)

	if err != nil {
		return fmt.Errorf("error loading aws config: %w", err)
	}

	client := s3.NewFromConfig(
		cfg,
		func(options *s3.Options) {
			options.UsePathStyle = true
		},
	)

	_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return fmt.Errorf("error creating s3 bucket: %w", err)
	}

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("file.txt"),
		Body:   strings.NewReader("lorem ipsum"),
	}, s3.WithAPIOptions())

	return nil
}
