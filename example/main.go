package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	crun "github.com/pip-services3-gox/pip-services3-commons-gox/run"
	ccount "github.com/pip-services3-gox/pip-services3-components-gox/count"
	clog "github.com/pip-services3-gox/pip-services3-components-gox/log"
	cservices "github.com/pip-services3-gox/pip-services3-rpc-gox/services"
	logic "github.com/pip-services3-gox/pip-services3-swagger-gox/example/logic"
	services "github.com/pip-services3-gox/pip-services3-swagger-gox/example/services"
	sservices "github.com/pip-services3-gox/pip-services3-swagger-gox/services"
)

func main() {
	ctx := context.Background()

	// Create components
	logger := clog.NewConsoleLogger()
	counter := ccount.NewLogCounters()
	controller := logic.NewDummyController()
	httpEndpoint := cservices.NewHttpEndpoint()
	restService := services.NewDummyRestService()
	httpService := services.NewDummyCommandableHttpService()
	statusService := cservices.NewStatusRestService()
	heartbeatService := cservices.NewHeartbeatRestService()
	swaggerService := sservices.NewSwaggerService()

	components := []any{
		logger,
		counter,
		controller,
		httpEndpoint,
		restService,
		httpService,
		statusService,
		heartbeatService,
		swaggerService,
	}

	// Configure components
	logger.Configure(ctx, cconf.NewConfigParamsFromTuples(
		"level", "trace",
	))

	httpEndpoint.Configure(ctx, cconf.NewConfigParamsFromTuples(
		"connection.prototol", "http",
		"connection.host", "localhost",
		"connection.port", 8080,
	))

	restService.Configure(ctx, cconf.NewConfigParamsFromTuples(
		"swagger.enable", true,
	))

	httpService.Configure(ctx, cconf.NewConfigParamsFromTuples(
		"base_route", "dummies2",
		"swagger.enable", true,
	))

	// Set references
	references := cref.NewReferencesFromTuples(ctx,
		cref.NewDescriptor("pip-services", "logger", "console", "default", "1.0"), logger,
		cref.NewDescriptor("pip-services", "counter", "log", "default", "1.0"), counter,
		cref.NewDescriptor("pip-services", "endpoint", "http", "default", "1.0"), httpEndpoint,
		cref.NewDescriptor("pip-services-dummies", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("pip-services-dummies", "service", "rest", "default", "1.0"), restService,
		cref.NewDescriptor("pip-services-dummies", "service", "commandable-http", "default", "1.0"), httpService,
		cref.NewDescriptor("pip-services", "status-service", "rest", "default", "1.0"), statusService,
		cref.NewDescriptor("pip-services", "heartbeat-service", "rest", "default", "1.0"), heartbeatService,
		cref.NewDescriptor("pip-services", "swagger-service", "http", "default", "1.0"), swaggerService,
	)

	cref.Referencer.SetReferences(ctx, references, components)

	// Open components
	err := crun.Opener.Open(ctx, "", components)
	if err != nil {
		logger.Error(ctx, "", err, "Failed to open components")
		return
	}

	// Wait until user presses ENTER
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press ENTER to stop the microservice...")
	reader.ReadString('\n')

	// Close components
	err = crun.Closer.Close(ctx, "", components)
	if err != nil {
		logger.Error(ctx, "", err, "Failed to close components")
	}
}
