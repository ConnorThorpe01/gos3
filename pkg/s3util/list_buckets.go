package s3util

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListBuckets(c *s3.Client) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	output, err := c.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	var names []string
	for _, b := range output.Buckets {
		if b.Name != nil {
			names = append(names, *b.Name)
		}
	}

	return names, nil
}
