package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
	"github.com/suhas-totality/gomongodemo/api/v1/service"
	"github.com/suhas-totality/gomongodemo/api/v1/service/rpc"
	"github.com/suhas-totality/gomongodemo/database/mongodb"
	"github.com/suhas-totality/gomongodemo/pkg/env"
	"github.com/suhas-totality/gomongodemo/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Init struct {
	envPlain *env.EnvPlain
	logger   *logger.Logger
}

func newInit(envPlain *env.EnvPlain) *Init {
	return &Init{
		envPlain: envPlain,
	}
}

func (i *Init) initLogger() *logger.Logger {
	i.logger = logger.NewLogger(i.envPlain.SERVER_LOG_LEVEL)
	defer i.logger.Log.Sync()
	i.logger.Log.Info("Server Boot; Initializing Services")
	return i.logger
}

func (i Init) initMongoDB(ctx context.Context) (*mongo.Client, func()) {
	mongodbURI := fmt.Sprintf(i.envPlain.MONGODB_URL, i.envPlain.MONGODB_USERNAME, i.envPlain.MONGODB_PASSWORD)

	mongodbClient, err := mongodb.NewMongoDBClient(ctx, mongodbURI)
	if err != nil {
		i.logger.Log.Fatal(err.Error())
	}

	err = mongodbClient.PingMongoClient()
	if err != nil {
		i.logger.Log.Fatal(err.Error())
	}

	log.Println("connected to MongoDB Server")

	return mongodbClient.Client, func() {
		mongodbClient.DisconnectMongoDB(ctx)
	}
}

func (i Init) initGRPCServer(serviceRegistry *service.ServiceRegistry) (func(), func()) {
	gRPCServer := NewServer(i.envPlain.SERVER_GRPC_PORT, time.Duration(i.envPlain.SERVER_TIMEOUT_SECONDS))
	pb.RegisterServiceServer(gRPCServer.GRPCServer, rpc.NewServer(serviceRegistry))
	return gRPCServer.StartServer, gRPCServer.StopServer
}
