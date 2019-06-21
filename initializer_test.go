package awsutil

import (
	"github.com/aphistic/sweet"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-nacelle/nacelle"
	. "github.com/onsi/gomega"
)

type InitializerSuite struct{}

func (s *InitializerSuite) TestInitDefault(t sweet.T) {
	factory := func(s *session.Session) interface{} {
		return "service"
	}

	initializer := NewServiceInitializer("test", factory)
	initializer.Logger = nacelle.NewNilLogger()
	initializer.Services = nacelle.NewServiceContainer()

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := initializer.Init(config)
	Expect(err).To(BeNil())
	Expect(initializer.Services.Get("test")).To(Equal("service"))
}

func (s *InitializerSuite) TestInitServiceSpecificConfig(t sweet.T) {
	var awsConfig *aws.Config
	factory := func(sess *session.Session) interface{} {
		awsConfig = sess.Config
		return "service"
	}

	initializer := NewServiceInitializer("test", factory)
	initializer.Logger = nacelle.NewNilLogger()
	initializer.Services = nacelle.NewServiceContainer()

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(map[string]string{
		"test_endpoint":   "http://localhost:1234",
		"test_log_level":  "debug",
		"test_region":     "local",
		"aws_disable_ssl": "true", // ignored
	}))

	err := initializer.Init(config)
	Expect(err).To(BeNil())
	Expect(initializer.Services.Get("test")).To(Equal("service"))
	Expect(*awsConfig.Endpoint).To(Equal("http://localhost:1234"))
	Expect(*awsConfig.LogLevel).To(Equal(aws.LogDebug))
	Expect(*awsConfig.Region).To(Equal("local"))
	Expect(*awsConfig.DisableSSL).To(BeFalse())
}

func (s *InitializerSuite) TestInitFallbackConfig(t sweet.T) {
	var awsConfig *aws.Config
	factory := func(sess *session.Session) interface{} {
		awsConfig = sess.Config
		return "service"
	}

	initializer := NewServiceInitializer("test", factory)
	initializer.Logger = nacelle.NewNilLogger()
	initializer.Services = nacelle.NewServiceContainer()

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(map[string]string{
		"aws_endpoint":    "http://localhost:1234",
		"aws_log_level":   "debug",
		"aws_region":      "local",
		"aws_disable_ssl": "true",
	}))

	err := initializer.Init(config)
	Expect(err).To(BeNil())
	Expect(initializer.Services.Get("test")).To(Equal("service"))
	Expect(*awsConfig.Endpoint).To(Equal("http://localhost:1234"))
	Expect(*awsConfig.LogLevel).To(Equal(aws.LogDebug))
	Expect(*awsConfig.Region).To(Equal("local"))
	Expect(*awsConfig.DisableSSL).To(BeTrue())
}

func (s *InitializerSuite) TestInitializeWithLiteralConfig(t sweet.T) {
	var awsConfig *aws.Config
	factory := func(sess *session.Session) interface{} {
		awsConfig = sess.Config
		return "service"
	}

	initializer := NewServiceInitializer("test", factory, &aws.Config{
		DisableSSL: aws.Bool(true),
	})

	initializer.Logger = nacelle.NewNilLogger()
	initializer.Services = nacelle.NewServiceContainer()

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(map[string]string{
		"aws_endpoint":  "http://localhost:1234",
		"aws_log_level": "debug",
		"aws_region":    "local",
	}))

	err := initializer.Init(config)
	Expect(err).To(BeNil())
	Expect(initializer.Services.Get("test")).To(Equal("service"))
	Expect(*awsConfig.Endpoint).To(Equal("http://localhost:1234"))
	Expect(*awsConfig.LogLevel).To(Equal(aws.LogDebug))
	Expect(*awsConfig.Region).To(Equal("local"))
	Expect(*awsConfig.DisableSSL).To(BeTrue())
}
