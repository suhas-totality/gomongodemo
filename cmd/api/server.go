package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	Port       int
	GRPCServer *grpc.Server
}

func NewServer(port int, connTimeout time.Duration) *Server {
	opts := []grpc.ServerOption{
		grpc.ConnectionTimeout(connTimeout * time.Second),
	}

	return &Server{
		Port:       port,
		GRPCServer: grpc.NewServer(opts...),
	}
}

func (s *Server) StartServer() {
	go func() {
		log.Println("grpc server listening on port", s.Port)
		listener, err := net.Listen("tcp", fmt.Sprint(":", s.Port))
		if err != nil {
			log.Fatal("unable to start gRPC server", err.Error())
		}
		if err := s.GRPCServer.Serve(listener); err != nil {
			log.Fatal("unable to register port", s.Port, err.Error())
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}

func (s *Server) StopServer() {
	fmt.Print("\b\b")
	s.GRPCServer.GracefulStop()
	log.Println("stopped grpc server; release port", s.Port)
}
