package config

import (
	"github.com/spf13/viper"

	"github.com/hdlproject/es-user-service/helper"
)

type (
	defaultConfig struct {
		configurable Configurable
	}
)

func newDefaultConfig(configurable Configurable) Configurable {
	viper.SetDefault("CONFIG_FILEPATH", ".")
	viper.SetDefault("CONFIG_FILENAME", ".env")

	viper.SetDefault("PORT", "7778")

	viper.SetDefault("DB_HOST", "127.0.0.1")
	viper.SetDefault("DB_PORT", "5433")
	viper.SetDefault("DB_NAME", "es-user-service")
	viper.SetDefault("DB_USER", "postgres")

	viper.SetDefault("EVENT_BUS_HOST", "127.0.0.1")
	viper.SetDefault("EVENT_BUS_PORT", "5672")
	viper.SetDefault("EVENT_BUS_USERNAME", "root")

	viper.SetDefault("EVENT_STORE_HOST", "127.0.0.1")
	viper.SetDefault("EVENT_STORE_PORT", "27018")
	viper.SetDefault("EVENT_STORE_USERNAME", "root")
	viper.SetDefault("EVENT_STORE_NAME", "es-user-service")

	viper.SetDefault("CENTRIFUGE_SERVER_URL", "ws://localhost:8000/connection/websocket")

	return defaultConfig{
		configurable: configurable,
	}
}

func (instance defaultConfig) Get() (config Config, err error) {
	config, err = getConfig()
	if err != nil {
		return Config{}, helper.WrapError(err)
	}

	return config, nil
}
