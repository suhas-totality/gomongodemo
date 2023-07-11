package rpc

import (
	"context"

	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
)

func (s *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	return s.Handler.CreateAccountHandler(ctx, req)
}
