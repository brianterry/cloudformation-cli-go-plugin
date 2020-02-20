package resource

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
)

func init() {
	err := os.Setenv("UNIT_TEST_TOGGLE", "TRUE")
	if err != nil {
		log.Println(err)
	}
}

func TestCreate(t *testing.T) {
	defer os.Unsetenv("UNIT_TEST_TOGGLE")

	tm := &Model{
		BucketName: aws.String("test"),
	}

	type args struct {
		req          handler.Request
		prevModel    *Model
		currentModel *Model
	}
	tests := []struct {
		name    string
		args    args
		want    handler.ProgressEvent
		wantErr bool
	}{
		{"Simple create test", args{handler.Request{}, nil, tm}, handler.ProgressEvent{
			ResourceModel:   tm,
			OperationStatus: handler.Success,
			Message:         "Completed",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.req, tt.args.prevModel, tt.args.currentModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
