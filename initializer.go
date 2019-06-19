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
	serviceTagModifiers := []nacelle.TagModifier{
		nacelle.NewFileTagSetter(),
		nacelle.NewEnvTagPrefixer(i.name),
		nacelle.NewFileTagPrefixer(i.name),
	}

	globalTagModifiers := []nacelle.TagModifier{
		nacelle.NewFileTagSetter(),
		nacelle.NewEnvTagPrefixer(i.name),
		nacelle.NewFileTagPrefixer(i.name),
	}

	serviceConfig := &Config{}
	if err := config.Load(&serviceConfig, serviceTagModifiers...); err != nil {
		return nil, err
	}

	if serviceConfig.IsDefault() {
		if err := config.Load(&serviceConfig, globalTagModifiers...); err != nil {
			return nil, err
		}
	}

	return serviceConfig.ToAWSConfig(i.Logger), nil
}
