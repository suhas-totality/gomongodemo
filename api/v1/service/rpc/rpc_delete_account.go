package rpc

import (
	"context"

	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
)

func (s *Server) DeleteAcount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	return s.Handler.DeleteAccountHandler(ctx, req)
}
