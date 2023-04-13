package file

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Volume struct {
	bucket string
	client *s3.Client
}

func (s S3Volume) Open(ctx context.Context, name string) (io.ReadCloser, error) {
	object, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(name),
	})
	if err != nil {
		return nil, err
	}

	return object.Body, nil
}

func NewS3Volume(ctx context.Context, region string, bucket string) (S3Volume, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return S3Volume{
		bucket: bucket,
		client: s3.NewFromConfig(cfg),
	}, nil
}

func NewS3VolumeCustomEndpoint(ctx context.Context, region string, bucket string, endpoint string) (S3Volume, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithEndpointResolverWithOptions(CustomEndpointResolver{endpoint: endpoint}),
		config.WithCredentialsProvider(aws.AnonymousCredentials{}),
	)

	if err != nil {
		return S3Volume{}, fmt.Errorf("error loading aws config: %w", err)
	}

	client := s3.NewFromConfig(
		cfg,
		func(options *s3.Options) {
			options.UsePathStyle = true
		},
	)

	return S3Volume{
		bucket: bucket,
		client: client,
	}, nil
}

type CustomEndpointResolver struct {
	endpoint string
}

func (c CustomEndpointResolver) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	return aws.Endpoint{
		PartitionID:   "aws",
		URL:           c.endpoint,
		SigningRegion: region,
	}, nil
}
