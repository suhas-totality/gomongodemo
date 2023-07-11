package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

func NewMongoDBClient(ctx context.Context, mongodbURI string) (*MongoDB, error) {

	mongodbClient, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return nil, fmt.Errorf("unable to create MongoDB Client %s", err.Error())
	}

	err = mongodbClient.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to establish connection with MongoDB Server %s", err.Error())
	}

	return &MongoDB{
		Client: mongodbClient,
	}, nil
}

func (s *MongoDB) DisconnectMongoDB(ctx context.Context) {
	s.Client.Disconnect(ctx)
}

func (s *MongoDB) PingMongoClient() error {
	return s.Client.Ping(context.TODO(), readpref.Primary())
}
