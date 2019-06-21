package awsutil

import (
	"github.com/aphistic/sweet"
	"github.com/go-nacelle/nacelle"
	. "github.com/onsi/gomega"
)

type ConfigSuite struct{}

func (s *ConfigSuite) TestILlegalLogLevel(t sweet.T) {
	config := &Config{RawLogLevel: "unknown"}
	Expect(config.PostLoad()).To(MatchError("unknown aws log level type 'unknown'"))
}

func (s *ConfigSuite) TestIsDefault(t sweet.T) {
	setterFuncs := []func(c *Config){
		func(c *Config) { c.CredentialsChainVerboseErrors = true },
		func(c *Config) { c.DisableComputeChecksums = true },
		func(c *Config) { c.DisableEndpointHostPrefix = true },
		func(c *Config) { c.DisableParamValidation = true },
		func(c *Config) { c.DisableRestProtocolURICleaning = true },
		func(c *Config) { c.DisableSSL = true },
		func(c *Config) { c.EC2MetadataDisableTimeoutOverride = true },
		func(c *Config) { c.EnableEndpointDiscovery = true },
		func(c *Config) { c.EnforceShouldRetryCheck = true },
		func(c *Config) { c.S3Disable100Continue = true },
		func(c *Config) { c.S3DisableContentMD5Validation = true },
		func(c *Config) { c.S3ForcePathStyle = true },
		func(c *Config) { c.S3UseAccelerate = true },
		func(c *Config) { c.UseDualStack = true },
		func(c *Config) { c.MaxRetries = 5 },
		func(c *Config) { c.Endpoint = "http://localhost:4569" },
		func(c *Config) { c.RawLogLevel = "debug" },
		func(c *Config) { c.Region = "us-east-1" },
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))

	for _, setterFunc := range setterFuncs {
		awsConfig := &Config{}
		err := config.Load(awsConfig)
		Expect(err).To(BeNil())

		Expect(awsConfig.IsDefault()).To(BeTrue())
		setterFunc(awsConfig)
		Expect(awsConfig.IsDefault()).To(BeFalse())
	}
}
