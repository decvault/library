package badger

const configSectionName = "badger"

type badgerConfig struct {
	Path string `yaml:"path"`
}
