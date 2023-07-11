package handler

import (
	"context"
	"strconv"

	"github.com/suhas-totality/gomongodemo/api/v1/models"
	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) ReadAccountHandler(ctx context.Context, req *pb.ReadAccountRequest) (res *pb.ReadAccountResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Error(codes.Internal, "ERROR_INTERNAL_SERVER")
		}
	}()

	h.CreateMaterializedView(ctx)

	db := h.serviceRegistry.MongoDBClient.Database("AccountsDB")
	coll := db.Collection("account_mat_view")

	var ppln bson.A
	var docFilters bson.A

	if req.AccountDetails.AccounNo.Accno <= 0 && req.AccountDetails.CustomerName.Name == "" && req.AccountDetails.Balance.Balance <= 0 {
		h.serviceRegistry.LoggerClient.Log.Warn("No Query Parameter is Entered, Please Enter atleast one Query Parameter")
	} else {
		if req.AccountDetails.AccounNo.Accno > 0 {
			docFilters = append(docFilters, bson.D{{Key: "text", Value: bson.D{{Key: "query", Value: strconv.Itoa(int(req.AccountDetails.AccounNo.Accno))}, {Key: "path", Value: "account_number"}}}})
		}
		if req.AccountDetails.CustomerName.Name != "" {
			docFilters = append(docFilters, bson.D{{Key: "text", Value: bson.D{{Key: "query", Value: req.AccountDetails.CustomerName.Name}, {Key: "path", Value: "name"}}}})
		}
		if req.AccountDetails.Balance.Balance > 0 {
			docFilters = append(docFilters, bson.D{{Key: "text", Value: bson.D{{Key: "query", Value: strconv.Itoa(int(req.AccountDetails.Balance.Balance))}, {Key: "path", Value: "balance"}}}})
		}

		compound := bson.D{{Key: "filter", Value: docFilters}}

		ppln = append(ppln, bson.D{{Key: "$search", Value: bson.D{{Key: "index", Value: "account_mat_search_index"}, {Key: "compound", Value: compound}}}})

		pipeline, err := coll.Aggregate(ctx, ppln)
		if err != nil {
			h.serviceRegistry.LoggerClient.Log.Fatal("Error in Aggregation Pipeline")
		}

		defer pipeline.Close(ctx)

		pbaccoll := &pb.AccountsCollection{}

		for pipeline.Next(ctx) {
			data := &models.AccountString{}
			err = pipeline.Decode(data)

			if err != nil {
				h.serviceRegistry.LoggerClient.Log.Error(ctx, err, "Error while retrieving data", data)
			}

			pbaccoll.Accounts = append(pbaccoll.Accounts, documentToAccount(data))
		}
		if err := pipeline.Err(); err != nil {
			h.serviceRegistry.LoggerClient.Log.Error(ctx, err, "Error while retrieving data")
		}

		return &pb.ReadAccountResponse{AccountsCollection: pbaccoll}, err
	}
	return nil, nil
}
