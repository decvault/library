package config

import (
	"github.com/pkg/errors"
	"go.uber.org/config"
	"go.uber.org/fx"
)

type AppConfigProvider struct {
	fx.Out

	Provider config.Provider
}

func newAppConfigProvider() (AppConfigProvider, error) {
	envProvider, err := newAppEnvConfigProvider()
	if err != nil {
		return AppConfigProvider{}, err
	}

	modeProvider, err := newAppModeConfigProvider()
	if err != nil {
		return AppConfigProvider{}, err
	}

	provider, err := config.NewProviderGroup("app", modeProvider, envProvider)
	if err != nil {
		return AppConfigProvider{}, errors.WithStack(err)
	}

	return AppConfigProvider{
		Provider: provider,
	}, nil
}

func Provide[TConfig any](configSectionName string) func(provider config.Provider) (TConfig, error) {
	return func(provider config.Provider) (TConfig, error) {
		var conf TConfig
		if err := provider.Get(configSectionName).Populate(&conf); err != nil {
			var empty TConfig
			return empty, errors.WithStack(err)
		}

		return conf, nil
	}
}
