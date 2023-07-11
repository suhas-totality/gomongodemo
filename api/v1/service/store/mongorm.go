package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Model struct {
}

func (m *Model) Create(ctx context.Context, db *mongo.Database, collection *mongo.Collection, model interface{}) error {
	_, err := collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) Read(ctx context.Context, db *mongo.Database, collection *mongo.Collection, filter interface{}, result interface{}) error {
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) Update(ctx context.Context, db *mongo.Database, collection *mongo.Collection, filter interface{}, update interface{}) error {
	ans, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if ans.MatchedCount == 0 {
		return status.Errorf(
			codes.NotFound,
			"Cannot find record to update",
		)
	}
	return nil
}

func (m *Model) Delete(ctx context.Context, db *mongo.Database, collection *mongo.Collection, filter interface{}) error {
	ans, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if ans.DeletedCount == 0 {
		return status.Errorf(
			codes.NotFound,
			"Cannot Find Record to delete",
		)
	}
	return nil
}
