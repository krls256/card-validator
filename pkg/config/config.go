package config

import (
	"errors"
	"fmt"
	"github.com/krls256/card-validator/pkg/transport/grpc"
	"github.com/krls256/card-validator/pkg/transport/http"
	"github.com/spf13/viper"
	"path/filepath"
)

var (
	ErrConfigParser = errors.New("config parser error")
)

type Config struct {
	HTTPConfig http.Config
	GRPCConfig grpc.Config
}

func New(path string) (*Config, error) {
	viper.Reset()

	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(filepath.Dir(abs))
	viper.SetConfigFile(filepath.Base(abs))

	config := &Config{}

	staticConfigs := map[string]interface{}{
		"grpc": &config.GRPCConfig,
		"http": &config.HTTPConfig,
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	for key, conf := range staticConfigs {
		if err := parseTagConfig(key, conf); err != nil {
			return nil, err
		}
	}

	return config, nil
}

func parseTagConfig(tag string, parseTo interface{}) error {
	subConfig := viper.Sub(tag)

	if err := parseSubConfig(subConfig, &parseTo, tag); err != nil {
		return err
	}

	return nil
}

func parseSubConfig(subConfig *viper.Viper, parseTo interface{}, name string) error {
	if subConfig == nil {
		return fmt.Errorf("%w: can not read %v config to %T: subconfig is nil", ErrConfigParser, name, parseTo)
	}

	if err := subConfig.Unmarshal(parseTo); err != nil {
		return err
	}

	return nil
}
