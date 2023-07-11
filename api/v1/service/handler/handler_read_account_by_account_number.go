package handler

import (
	"context"

	"github.com/suhas-totality/gomongodemo/api/v1/models"
	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) ReadAccountByAccountNumberHandler(ctx context.Context, req *pb.ReadAccountByAccountNumberRequest) (res *pb.ReadAccountByAccountNumberResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Error(codes.Internal, "ERROR_INTERNAL_SERVER")
		}
	}()

	db := h.serviceRegistry.MongoDBClient.Database("AccountsDB")
	coll := db.Collection("AccountsCollection")
	var acc *models.Account

	acc, err = h.store.Account.ReadAccountByAccountNumber(ctx, db, coll, &models.AccNumber{
		Accno: req.AccounNo.Accno,
	})
	if err != nil {
		return nil, nil
	}

	pbac := &pb.Account{AccounNo: &pb.AccNumber{Accno: acc.AccNumber},
		CustomerName: &pb.Name{Name: acc.Name},
		Balance:      &pb.Balance{Balance: acc.Balance}}

	return &pb.ReadAccountByAccountNumberResponse{AccountDetails: pbac}, err
}
