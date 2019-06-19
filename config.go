package awsutil

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-nacelle/nacelle"
)

type Config struct {
	CredentialsChainVerboseErrors     bool   `env:"credentials_chain_verbose_errors"`
	DisableComputeChecksums           bool   `env:"disable_compute_checksums"`
	DisableEndpointHostPrefix         bool   `env:"disable_endpoint_host_prefix"`
	DisableParamValidation            bool   `env:"disable_param_validation"`
	DisableRestProtocolURICleaning    bool   `env:"disable_rest_protocol_uri_cleaning"`
	DisableSSL                        bool   `env:"disable_ssl"`
	EC2MetadataDisableTimeoutOverride bool   `env:"ec2_metadata_disable_timeout_override"`
	EnableEndpointDiscovery           bool   `env:"enable_endpoint_discovery"`
	Endpoint                          string `env:"endpoint"`
	EnforceShouldRetryCheck           bool   `env:"enforce_should_retry_check"`
	MaxRetries                        int    `env:"max_retries" default:"-1"`
	RawLogLevel                       string `env:"log_level"`
	Region                            string `env:"region"`
	S3Disable100Continue              bool   `env:"s3_disable_100_continue"`
	S3DisableContentMD5Validation     bool   `env:"s3_disable_content_md5_validation"`
	S3ForcePathStyle                  bool   `env:"s3_force_path_style"`
	S3UseAccelerate                   bool   `env:"s3_useaccelerate"`
	UseDualStack                      bool   `env:"use_dual_stack"`
	LogLevel                          aws.LogLevelType
}

var levelMap = map[string]aws.LogLevelType{
	"off":                          aws.LogOff,
	"debug":                        aws.LogDebug,
	"debug_with_signing":           aws.LogDebugWithSigning,
	"debug_with_http_body":         aws.LogDebugWithHTTPBody,
	"debug_with_request_retries":   aws.LogDebugWithRequestRetries,
	"debug_with_request_errors":    aws.LogDebugWithRequestErrors,
	"debug_with_event_stream_body": aws.LogDebugWithEventStreamBody,
}

func (c *Config) PostLoad() error {
	level, ok := levelMap[strings.ToLower(c.RawLogLevel)]
	if !ok {
		return fmt.Errorf("unknown aws log level type '%s'", c.RawLogLevel)
	}

	c.LogLevel = level
	return nil
}

func (c *Config) IsDefault() bool {
	if c.Endpoint != "" || c.MaxRetries != -1 || c.RawLogLevel != "off" || c.Region != "" {
		return false
	}

	bools := []bool{
		c.CredentialsChainVerboseErrors,
		c.DisableComputeChecksums,
		c.DisableEndpointHostPrefix,
		c.DisableParamValidation,
		c.DisableRestProtocolURICleaning,
		c.DisableSSL,
		c.EC2MetadataDisableTimeoutOverride,
		c.EnableEndpointDiscovery,
		c.EnforceShouldRetryCheck,
		c.S3Disable100Continue,
		c.S3DisableContentMD5Validation,
		c.S3ForcePathStyle,
		c.S3UseAccelerate,
		c.UseDualStack,
	}

	for _, b := range bools {
		if b {
			return false
		}
	}

	return true
}

func (c *Config) ToAWSConfig(logger nacelle.Logger) *aws.Config {
	return &aws.Config{
		CredentialsChainVerboseErrors:     aws.Bool(c.CredentialsChainVerboseErrors),
		DisableComputeChecksums:           aws.Bool(c.DisableComputeChecksums),
		DisableEndpointHostPrefix:         aws.Bool(c.DisableEndpointHostPrefix),
		DisableParamValidation:            aws.Bool(c.DisableParamValidation),
		DisableRestProtocolURICleaning:    aws.Bool(c.DisableRestProtocolURICleaning),
		DisableSSL:                        aws.Bool(c.DisableSSL),
		EC2MetadataDisableTimeoutOverride: aws.Bool(c.EC2MetadataDisableTimeoutOverride),
		EnableEndpointDiscovery:           aws.Bool(c.EnableEndpointDiscovery),
		Endpoint:                          aws.String(c.Endpoint),
		EnforceShouldRetryCheck:           aws.Bool(c.EnforceShouldRetryCheck),
		Logger:                            NewAWSLogAdapter(logger),
		LogLevel:                          aws.LogLevel(c.LogLevel),
		MaxRetries:                        aws.Int(c.MaxRetries),
		Region:                            aws.String(c.Region),
		S3Disable100Continue:              aws.Bool(c.S3Disable100Continue),
		S3DisableContentMD5Validation:     aws.Bool(c.S3DisableContentMD5Validation),
		S3ForcePathStyle:                  aws.Bool(c.S3ForcePathStyle),
		S3UseAccelerate:                   aws.Bool(c.S3UseAccelerate),
		UseDualStack:                      aws.Bool(c.UseDualStack),
	}
}
