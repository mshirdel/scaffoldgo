package config

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Server  Server  `mapstructure:"server" validate:"required"`
	Logging Logging `mapstructure:"logging" validate:"required"`
}

type Server struct {
	Address      string        `mapstructure:"address" validate:"required,hostname_port"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout" validate:"required"`
	WriteTimeout time.Duration `mapstructure:"write-timeout" validate:"required"`
	IdleTimeout  time.Duration `mapstructure:"idle-timeout" validate:"required"`
}

type Logging struct {
	Level string `mapstructure:"level" validate:"required,lowercase,oneof=trace debug info warn warning error fatal panic"`
}

const _envPrefix = "{{.ModuleName}}"

func InitViper(configPath string) (*Config, error) {
	var c Config

	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewReader([]byte(_builtinConfig))); err != nil {
		return nil, fmt.Errorf("error loading default configs: %w", err)
	}

	v.SetConfigFile(configPath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.SetEnvPrefix(_envPrefix)
	v.AutomaticEnv()

	switch err := v.MergeInConfig(); err.(type) {
	case nil:
	case *os.PathError:
		logrus.Infof("config file (%s) not found, Using defaults and environment variables", configPath)
	default:
		logrus.Warnf("failed to load config file: %s", err)
	}

	if err := v.UnmarshalExact(&c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config into struct: %w", err)
	}

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &c, nil
}

func (c *Config) Validate() error {
	return validator.New().Struct(c)
}
