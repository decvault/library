package grpcsrv

import (
	"github.com/decvault/library/common/config"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"grpcsrv",
		fx.Provide(
			newGrpcServer,
			config.Provide[grpcServerConfig](configSectionName),
		),
	)
}
