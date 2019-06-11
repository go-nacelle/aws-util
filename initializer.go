package awsutil

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-nacelle/nacelle"
)

type (
	ServiceInitializer struct {
		Logger   nacelle.Logger           `service:"logger"`
		Services nacelle.ServiceContainer `service:"services"`
		name     string
		factory  Factory
		configs  []*aws.Config
	}

	Factory func(sess *session.Session) interface{}
)

func NewServiceInitializer(name string, factory Factory, configs ...*aws.Config) nacelle.Initializer {
	return &ServiceInitializer{
		name:    name,
		factory: factory,
		configs: configs,
	}
}

func (i *ServiceInitializer) Init(config nacelle.Config) error {
	awsConfig, err := i.loadConfig(config)
	if err != nil {
		return err
	}

	sess, err := session.NewSession(append(
		[]*aws.Config{awsConfig},
		i.configs...,
	)...)

	if err != nil {
		return err
	}

	return i.Services.Set(i.name, i.factory(sess))
}

func (i *ServiceInitializer) loadConfig(config nacelle.Config) (*aws.Config, error) {
	tagModifiers := []TagModifier{
		NewFileTagSetter(),
		NewEnvTagPrefixer(i.name),
		NewFileTagPrefixer(i.name),
	}

	serviceConfig := &Config{}
	if err := config.Load(&serviceConfig, tagModifiers...); err != nil {
		return nil, err
	}

	if serviceConfig.IsDefault() {
		if err := config.Load(&serviceConfig, tagModifiers[0]); err != nil {
			return nil, err
		}
	}

	return serviceConfig.ToAWSConfig(i.Logger), nil
}
