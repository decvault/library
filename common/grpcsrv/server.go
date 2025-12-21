package grpcsrv

import (
	"context"
	"fmt"
	"net"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GrpcServer interface {
	Run(context.Context) error
	GracefulStop(context.Context)
}

type grpcServer struct {
	server   *grpc.Server
	listener net.Listener

	logger *logrus.Logger
}

func newGrpcServer(
	config grpcServerConfig,
	opts SetupOpts,
	setupServer SetupFunc,
	logger *logrus.Logger,
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
		logger:   logger,
	}, nil
}

func (s *grpcServer) Run(ctx context.Context) error {
	s.logger.WithContext(ctx).Debugf("starting gprc server")
	if err := s.server.Serve(s.listener); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *grpcServer) GracefulStop(ctx context.Context) {
	s.logger.WithContext(ctx).Debugf("stopping gprc server gracefully")
	s.server.GracefulStop()
	s.logger.WithContext(ctx).Debugf("grpc server stopped gracefully")
}
