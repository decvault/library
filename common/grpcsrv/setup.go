package grpcsrv

import "google.golang.org/grpc"

type SetupOpts = []grpc.ServerOption

type SetupFunc func(*grpc.Server)
