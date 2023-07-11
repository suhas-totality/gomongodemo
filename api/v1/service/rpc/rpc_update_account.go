package rpc

import (
	"context"

	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
)

func (s *Server) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	return s.Handler.UpdateAccountHandler(ctx, req)
}
