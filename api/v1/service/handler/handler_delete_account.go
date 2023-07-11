package handler

import (
	"context"

	"github.com/suhas-totality/gomongodemo/api/v1/models"
	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) DeleteAccountHandler(ctx context.Context, req *pb.DeleteAccountRequest) (res *pb.DeleteAccountResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Error(codes.Internal, "ERROR_INTERNAL_SERVER")
		}
	}()

	db := h.serviceRegistry.MongoDBClient.Database("AccountsDB")
	coll := db.Collection("AccountsCollection")
	matvw := db.Collection("account_mat_view")

	ok1, err := h.store.Account.DeleteAccountByAccountNumber(ctx, db, coll, &models.AccNumber{
		Accno: req.AccounNo.Accno,
	})
	if err != nil {
		return nil, nil
	}

	ok2, err := h.store.Account.DeleteAccountByAccountNumberMatView(ctx, db, matvw, &models.AccNumber{
		Accno: req.AccounNo.Accno,
	})
	if err != nil {
		return nil, nil
	}
	return &pb.DeleteAccountResponse{DeleteAccountSuccess: ok1 && ok2}, err
}
