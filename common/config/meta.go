package config

import (
	"os"

	"github.com/pkg/errors"
)

type AppMeta struct {
	Stage Stage
}

func newAppMeta() (AppMeta, error) {
	stage := Stage(os.Getenv("APP_STAGE"))
	if stage != Dev && stage != Test && stage != Prod {
		return AppMeta{}, errors.Errorf("unknown application stage: %s", stage)
	}

	return AppMeta{
		Stage: stage,
	}, nil
}
