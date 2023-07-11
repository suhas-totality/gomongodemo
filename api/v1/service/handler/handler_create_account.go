package handler

import (
	"context"

	"github.com/suhas-totality/gomongodemo/api/v1/models"
	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) CreateAccountHandler(ctx context.Context, req *pb.CreateAccountRequest) (res *pb.CreateAccountResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Error(codes.Internal, "ERROR_INTERNAL_SERVER")
		}
	}()

	db := h.serviceRegistry.MongoDBClient.Database("AccountsDB")
	coll := db.Collection("AccountsCollection")

	ok, err := h.store.Account.CreateNewAccount(ctx, db, coll, &models.Account{
		AccNumber: req.AccountDetails.GetAccounNo().Accno,
		Name:      req.AccountDetails.GetCustomerName().GetName(),
		Balance:   req.GetAccountDetails().GetBalance().Balance,
	})
	if err != nil {
		return nil, nil
	}
	h.CreateMaterializedView(ctx)
	return &pb.CreateAccountResponse{CreateAccountSuccess: ok}, err
}
