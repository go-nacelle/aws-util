package awsutil

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-nacelle/nacelle"
)

type AWSLogAdapter struct {
	logger nacelle.Logger
}

func NewAWSLogAdapter(logger nacelle.Logger) aws.Logger {
	return &AWSLogAdapter{
		logger: logger,
	}
}

func (a *AWSLogAdapter) Log(args ...interface{}) {
	if len(args) == 0 {
		return
	}

	format, ok := args[0].(string)
	if !ok {
		format = fmt.Sprintf("%s", args[0])
	}

	a.logger.Debug(format, args[1:]...)
}
