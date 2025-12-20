package badger

import (
	"context"

	"github.com/dgraph-io/badger/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func newDB(
	opts *badger.Options,
	lc fx.Lifecycle,
) (*badger.DB, error) {
	db, err := badger.Open(*opts)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return errors.WithStack(db.Close())
		},
	})

	return db, nil
}

func newOpts(
	config badgerConfig,
	logger *logrus.Logger,
) *badger.Options {
	opts := badger.
		DefaultOptions(config.Path).
		WithLogger(logger)

	return &opts
}
