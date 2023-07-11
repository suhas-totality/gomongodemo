package rpc

import (
	"context"

	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
)

func (s *Server) ReadAccount(ctx context.Context, req *pb.ReadAccountRequest) (*pb.ReadAccountResponse, error) {
	return s.Handler.ReadAccountHandler(ctx, req)
}
