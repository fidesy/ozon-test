package grpc

import (
	"net"

	"github.com/fidesy/ozon-test/internal/service"
	shortener "github.com/fidesy/ozon-test/proto"

	"google.golang.org/grpc"
)

type Server struct {
	shortener.UnimplementedURLServiceServer
	service *service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) Start(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	shortener.RegisterURLServiceServer(grpcServer, s)

	return grpcServer.Serve(lis)
}
