package main

import (
	"github.com/go-nacelle/awsutil"
	"github.com/go-nacelle/httpbase"
	"github.com/go-nacelle/nacelle"
)

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
	processes.RegisterInitializer(awsutil.NewDynamoDBServiceInitializer(), nacelle.WithInitializerName("dynamodb"))
	processes.RegisterProcess(httpbase.NewServer(NewServerInitializer()), nacelle.WithProcessName("http-server"))
	return nil
}

func main() {
	nacelle.NewBootstrapper("awsutil-example", setup).BootAndExit()
}
