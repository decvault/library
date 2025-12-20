package badger

import (
	"github.com/decvault/library/common/config"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"badger",
		fx.Provide(
			newDB,
			newOpts,
			config.Provide[badgerConfig](configSectionName),
		),
	)
}
