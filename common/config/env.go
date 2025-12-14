package config

import (
	"os"

	"github.com/pkg/errors"
	"go.uber.org/config"
)

const (
	envConfigTemplateFilePath = "config/config.env.yaml"
)

func newAppEnvConfigProvider() (returnConf *config.YAML, returnErr error) {
	cfg, err := config.NewYAML(
		config.File(envConfigTemplateFilePath),
		config.Expand(os.LookupEnv),
	)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return cfg, nil
}
