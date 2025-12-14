package config

import (
	"github.com/pkg/errors"
	"go.uber.org/config"
)

type Stage string

const (
	Dev  Stage = "development"
	Test Stage = "testing"
	Prod Stage = "production"
)

const (
	devConfigFilePath  = "config/config.dev.yaml"
	testConfigFilePath = "config/config.test.yaml"
	prodConfigFilePath = "config/config.prod.yaml"
)

func newAppStageConfigProvider(meta AppMeta) (*config.YAML, error) {
	configFilePath, err := getAppStageConfigFilePath(meta)
	if err != nil {
		return nil, err
	}

	yaml, err := config.NewYAML(config.File(configFilePath))
	return yaml, errors.WithStack(err)
}

func getAppStageConfigFilePath(meta AppMeta) (string, error) {
	switch meta.Stage {
	case Dev:
		return devConfigFilePath, nil
	case Test:
		return testConfigFilePath, nil
	case Prod:
		return prodConfigFilePath, nil
	default:
		return "", errors.Errorf("unknown application stage: %s", meta.Stage)
	}
}
