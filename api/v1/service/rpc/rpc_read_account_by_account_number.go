package rpc

import (
	"context"

	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
)

func (s *Server) ReadAccountByAccountNumber(ctx context.Context, req *pb.ReadAccountByAccountNumberRequest) (*pb.ReadAccountByAccountNumberResponse, error) {
	return s.Handler.ReadAccountByAccountNumberHandler(ctx, req)
}
