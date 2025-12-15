package logging

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func newLogger(config loggerConfig) (*logrus.Logger, error) {
	logger := logrus.New()

	err := configureLogger(logger, config)
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func configureLogger(logger *logrus.Logger, config loggerConfig) error {
	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return errors.WithStack(err)
	}

	logger.SetLevel(logLevel)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
		PadLevelText:    true,
	})

	logger.AddHook(&ContextHook{})

	return nil
}
