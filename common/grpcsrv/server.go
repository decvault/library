package grpcsrv

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GrpcServer interface {
	Run() error
}

type grpcServer struct {
	server   *grpc.Server
	listener net.Listener
}

func newGrpcServer(
	config grpcServerConfig,
	opts SetupOpts,
	setupServer SetupFunc,
) (GrpcServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	server := grpc.NewServer(opts...)

	setupServer(server)

	return &grpcServer{
		server:   server,
		listener: listener,
	}, nil
}

func (s *grpcServer) Run() error {
	err := s.server.Serve(s.listener)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
