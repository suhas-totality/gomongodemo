package handler

import (
	"context"

	"github.com/suhas-totality/gomongodemo/api/v1/models"
	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) UpdateAccountHandler(ctx context.Context, req *pb.UpdateAccountRequest) (res *pb.UpdateAccountResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Error(codes.Internal, "ERROR_INTERNAL_SERVER")
		}
	}()

	db := h.serviceRegistry.MongoDBClient.Database("AccountsDB")
	coll := db.Collection("AccountsCollection")

	ok, err := h.store.Account.UpdateAccountByAccountNumber(ctx, db, coll, &models.Account{
		AccNumber: req.AccountDetails.GetAccounNo().Accno,
		Name:      req.AccountDetails.GetCustomerName().GetName(),
		Balance:   req.GetAccountDetails().GetBalance().Balance,
	})
	if err != nil {
		return nil, nil
	}

	h.CreateMaterializedView(ctx)
	return &pb.UpdateAccountResponse{UpdateAccountSuccess: ok}, err
}
