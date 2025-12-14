package logging

import (
	"github.com/decvault/library/common/config"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"logging",
		fx.Provide(
			newLogger,
			config.Provide[loggerConfig](configSectionName),
		),
	)
}
