package config

import (
	"errors"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func Load() (*Config, error) {
	path := getConfigPath(os.Getenv("APP_CONFIG"))
	v, err := LoadConfig(path, "yaml")
	if err != nil {
		return nil, err
	}
	cfg, err := ParseConfig(v)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.AllowEmptyEnv(true)

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, err
		}
	}

	return v, nil
}

func getConfigPath(env string) string {
	if env == "production" {
		return "/config/production"
	} else {
		return "/config/testing"
	}
}
