package resource

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// S3Mock is used to mock calls to s3.
// For more information on how to mock the AWS Go SDK, view our blog
// https://aws.amazon.com/blogs/developer/mocking-out-then-aws-sdk-for-go-for-unit-testing/
type S3Mock struct {
	s3iface.S3API
}

func newMockS3() *S3Mock {
	return &S3Mock{}
}

// CreateBucket mocks the S3 CreateBucket method.
func (q *S3Mock) CreateBucket(*s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	return nil, nil
}
