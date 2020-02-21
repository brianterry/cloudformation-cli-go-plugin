package resource

import (
	"log"
	"os"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	var sess *session.Session
	var svc s3iface.S3API

	if os.Getenv("UNIT_TEST_TOGGLE") == "TRUE" {
		// Use the test type.
		svc = newMockS3()

	} else {
		// Get Session from the request.
		sess = req.Session
		svc = s3.New(sess)
	}
	// Create the S3 Bucket
	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: currentModel.BucketName,
	})

	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
		}, nil
	}
	log.Printf("Bucket %s successfully created\n", *currentModel.BucketName)
	// Return a new ProgressEvent
	p := handler.NewProgressEvent()
	p.ResourceModel = currentModel
	p.OperationStatus = handler.Success
	p.Message = "Completed"

	// return the status
	return p, nil

}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	var sess *session.Session
	var svc s3iface.S3API

	if os.Getenv("UNIT_TEST_TOGGLE") == "TRUE" {
		// Use the test type.
		svc = newMockS3()

	} else {
		// Get Session from the request.
		sess = req.Session
		svc = s3.New(sess)
	}

	// List the S3 Bucket
	buckets, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
		}, nil
	}
	for _, bucket := range buckets.Buckets {

		if *bucket.Name == *currentModel.BucketName {
			log.Printf("Bucket %s successfully found\n", *currentModel.BucketName)
			// Return a new ProgressEvent
			p := handler.NewProgressEvent()
			p.ResourceModel = currentModel
			p.OperationStatus = handler.Success
			p.Message = "Completed"
		}
	}

	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          "Resource not found",
		HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Because we only have one property, and it's the primaryIdentifier, to change such a property on a live resource,
	// we replace that resource by creating a new instance of the resource and terminating the old one.
	// Cloudformation handles the delete, so we just call the Create function.

	return Create(req, prevModel, currentModel)
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	var sess *session.Session
	var svc s3iface.S3API

	if os.Getenv("UNIT_TEST_TOGGLE") == "TRUE" {
		// Use the test type.
		svc = newMockS3()

	} else {
		// Get Session from the request.
		sess = req.Session
		svc = s3.New(sess)
	}

	// Delete the S3 Bucket
	_, err := svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: currentModel.BucketName,
	})

	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
		}, nil
	}
	log.Printf("Bucket %s successfully deleted\n", *currentModel.BucketName)

	p := handler.NewProgressEvent()
	p.ResourceModel = currentModel
	p.OperationStatus = handler.Success
	p.Message = "Completed"

	// return the status
	return p, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	var sess *session.Session
	var svc s3iface.S3API
	var models []interface{}

	if len(*currentModel.BucketName) == 0 {
		currentModel.BucketName = aws.String(req.LogicalResourceID)
	}
	if os.Getenv("UNIT_TEST_TOGGLE") == "TRUE" {
		// Use the test type.
		svc = newMockS3()

	} else {
		// Get Session from the request.
		sess = req.Session
		svc = s3.New(sess)
	}

	// List the S3 Bucket
	// we are returning all the buckets, but this may not always be need.
	buckets, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
		}, nil
	}
	for _, bucket := range buckets.Buckets {

		models = append(models, &Model{
			BucketName: bucket.Name,
		})

		if *bucket.Name == *currentModel.BucketName {
			log.Printf("Bucket %s successfully found\n", *currentModel.BucketName)
			// Return a new ProgressEvent
			p := handler.NewProgressEvent()
			p.ResourceModels = models
			p.OperationStatus = handler.Success
			p.Message = "Completed"
		}
	}

	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          "Resource not found",
		HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
	}, nil
}
