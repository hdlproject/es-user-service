package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"

	"github.com/hdlproject/es-user-service/helper"
)

type (
	Config struct {
		Port         string
		Database     Database
		EventBus     EventBus
		EventStorage EventStorage
		AWS          AWS
		Centrifuge   Centrifuge
	}

	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		Name     string
	}

	EventBus struct {
		Host     string
		Port     string
		Username string
		Password string
	}

	EventStorage struct {
		Host     string
		Port     string
		Username string
		Password string
		Name     string
	}

	AWS struct {
		ID     string
		Secret string
	}

	Centrifuge struct {
		ServerUrl string
		Token     string
	}
)

const (
	missingConfigError = "config %s is missing"
)

var (
	instance *Config
)

func GetInstance() (Config, error) {
	configurable, err := newConfig()
	if err != nil {
		return Config{}, helper.WrapError(err)
	}

	if instance == nil {
		config, err := configurable.Get()
		if err != nil {
			return Config{}, helper.WrapError(err)
		}

		instance = &config
	}

	return *instance, nil
}

func newConfig() (Configurable, error) {
	defaultConfigurable := newDefaultConfig(nil)

	envConfigurable, err := newEnvConfig(defaultConfigurable)
	if err != nil {
		return nil, helper.WrapError(err)
	}

	fileConfigurable, err := newFileConfig(envConfigurable)
	if err != nil {
		return nil, helper.WrapError(err)
	}

	return fileConfigurable, nil
}

func getConfig() (Config, error) {
	dbPassword, err := getMandatoryString("DB_PASSWORD")
	if err != nil {
		return Config{}, helper.WrapError(err)
	}

	eventBusPassword, err := getMandatoryString("EVENT_BUS_PASSWORD")
	if err != nil {
		return Config{}, helper.WrapError(err)
	}

	eventStoragePassword, err := getMandatoryString("EVENT_STORAGE_PASSWORD")
	if err != nil {
		return Config{}, helper.WrapError(err)
	}

	return Config{
		Port: viper.GetString("PORT"),
		Database: Database{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			Username: viper.GetString("DB_USERNAME"),
			Password: dbPassword,
			Name:     viper.GetString("DB_NAME"),
		},
		EventBus: EventBus{
			Host:     viper.GetString("EVENT_BUS_HOST"),
			Port:     viper.GetString("EVENT_BUS_PORT"),
			Username: viper.GetString("EVENT_BUS_USERNAME"),
			Password: eventBusPassword,
		},
		EventStorage: EventStorage{
			Host:     viper.GetString("EVENT_STORAGE_HOST"),
			Port:     viper.GetString("EVENT_STORAGE_PORT"),
			Username: viper.GetString("EVENT_STORAGE_USERNAME"),
			Password: eventStoragePassword,
			Name:     viper.GetString("EVENT_STORAGE_NAME"),
		},
		AWS: AWS{
			ID:     viper.GetString("AWS_ID"),
			Secret: viper.GetString("AWS_SECRET"),
		},
		Centrifuge: Centrifuge{
			ServerUrl: viper.GetString("CENTRIFUGE_SERVER_URL"),
			Token:     viper.GetString("CENTRIFUGE_TOKEN"),
		},
	}, nil
}

func getMandatoryString(key string) (string, error) {
	if !viper.IsSet(key) {
		return "", helper.WrapError(errors.New(fmt.Sprintf(missingConfigError, key)))
	}

	value := viper.GetString(key)
	return value, nil
}
