package handler

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (h Handler) CreateMaterializedView(ctx context.Context) {
	db := h.serviceRegistry.MongoDBClient.Database("AccountsDB")
	coll := db.Collection("AccountsCollection")

	pipeln, err := coll.Aggregate(ctx,
		bson.A{
			bson.D{
				{Key: "$project",
					Value: bson.D{
						{Key: "account_number", Value: bson.D{{Key: "$toString", Value: "$account_number"}}},
						{Key: "name", Value: "$name"},
						{Key: "balance", Value: bson.D{{Key: "$toString", Value: "$balance"}}},
					},
				},
			},
			bson.D{
				{Key: "$merge",
					Value: bson.D{
						{Key: "into", Value: "account_mat_view"},
						{Key: "whenMatched", Value: "replace"},
					},
				},
			},
		})
	if err != nil {
		h.serviceRegistry.LoggerClient.Log.Fatal("Error while creating Materialized View")
	}

	defer pipeln.Close(ctx)
}
