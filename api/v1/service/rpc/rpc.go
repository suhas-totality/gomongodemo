package rpc

import (
	"github.com/suhas-totality/gomongodemo/api/v1/proto/pb"
	"github.com/suhas-totality/gomongodemo/api/v1/service"
	"github.com/suhas-totality/gomongodemo/api/v1/service/handler"
)

type Server struct {
	Handler handler.Handler
	pb.ServiceServer
}

func NewServer(svc *service.ServiceRegistry) *Server {
	return &Server{
		Handler: handler.NewHandler(svc),
	}
}
