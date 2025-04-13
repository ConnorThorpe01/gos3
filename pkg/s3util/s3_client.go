package s3util

import (
	"fmt"

	"github.com/ConnorThorpe01/gos3/pkg/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client(cfg *config.Config) (*s3.Client, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           fmt.Sprintf("http://%s", cfg.S3.Endpoint),
			SigningRegion: cfg.S3.Region, // Use the region from the config
		}, nil
	})

	awsCfg := aws.Config{
		Credentials:                 aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(cfg.S3.AccessKey, cfg.S3.SecretKey, "")),
		EndpointResolverWithOptions: customResolver,
		Region:                      cfg.S3.Region,
	}

	client := s3.NewFromConfig(awsCfg)
	return client, nil
}
