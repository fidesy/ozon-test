package grpc

import (
	"net"

	"github.com/fidesy/ozon-test/internal/usecase"
	shortener "github.com/fidesy/ozon-test/proto"

	"google.golang.org/grpc"
)

type Server struct {
	shortener.UnimplementedURLServiceServer
	usecases *usecase.Usecase
}

func NewServer(usecases *usecase.Usecase) *Server {
	return &Server{
		usecases: usecases,
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
