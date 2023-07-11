package handler

import (
	"strconv"

	"github.com/suhas-totality/gomongodemo/api/v1/models"
	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
)

func documentToAccount(data *models.AccountString) *pb.Account {
	acno, _ := strconv.Atoi(data.AccNumber)
	bal, _ := strconv.ParseFloat(data.Balance, 32)
	return &pb.Account{
		AccounNo:     &pb.AccNumber{Accno: int32(acno)},
		CustomerName: &pb.Name{Name: data.Name},
		Balance:      &pb.Balance{Balance: bal},
	}
}
