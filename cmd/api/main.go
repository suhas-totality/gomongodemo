package main

import (
	"context"

	"github.com/suhas-totality/gomongodemo/api/v1/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var releaseMongoDBClient func()

	serviceRegistry := service.NewServiceRegistry()
	serviceRegistry.LoggerClient = preInit.initLogger()
	serviceRegistry.MongoDBClient, releaseMongoDBClient = preInit.initMongoDB(ctx)

	startGRPCServer, releaseGRPCServer := preInit.initGRPCServer(serviceRegistry)

	defer releaseMongoDBClient()
	defer releaseGRPCServer()
	defer cancel()

	serviceRegistry.LoggerClient.Log.Info("Services Initialized; Starting Server")
	startGRPCServer()
}
