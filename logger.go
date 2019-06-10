package awsutil

import (
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
	a.logger.Debug("AWS: %#v", args) // TODO
}
