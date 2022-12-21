package awsutil

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-nacelle/nacelle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitDefault(t *testing.T) {
	factory := func(s *session.Session) interface{} {
		return "service"
	}

	initializer := NewServiceInitializer("test", factory)
	initializer.Logger = nacelle.NewNilLogger()
	initializer.Services = nacelle.NewServiceContainer()

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := initializer.Init(config)
	require.Nil(t, err)

	service, err := initializer.Services.Get("test")
	require.Nil(t, err)

	assert.Equal(t, service, "service")
}

func TestInitServiceSpecificConfig(t *testing.T) {
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
	require.Nil(t, err)

	service, err := initializer.Services.Get("test")
	require.Nil(t, err)

	assert.Equal(t, service, "service")
	assert.Equal(t, *awsConfig.Endpoint, "http://localhost:1234")
	assert.Equal(t, *awsConfig.LogLevel, aws.LogDebug)
	assert.Equal(t, *awsConfig.Region, "local")
	assert.False(t, *awsConfig.DisableSSL)
}

func TestInitFallbackConfig(t *testing.T) {
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
	require.Nil(t, err)

	service, err := initializer.Services.Get("test")
	require.Nil(t, err)

	assert.Equal(t, service, "service")
	assert.Equal(t, *awsConfig.Endpoint, "http://localhost:1234")
	assert.Equal(t, *awsConfig.LogLevel, aws.LogDebug)
	assert.Equal(t, *awsConfig.Region, "local")
	assert.True(t, *awsConfig.DisableSSL)
}

func TestInitializeWithLiteralConfig(t *testing.T) {
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
	require.Nil(t, err)

	service, err := initializer.Services.Get("test")
	require.Nil(t, err)

	assert.Equal(t, service, "service")
	assert.Equal(t, *awsConfig.Endpoint, "http://localhost:1234")
	assert.Equal(t, *awsConfig.LogLevel, aws.LogDebug)
	assert.Equal(t, *awsConfig.Region, "local")
	assert.True(t, *awsConfig.DisableSSL)
}
