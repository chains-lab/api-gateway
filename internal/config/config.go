package config

import (
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
}

type LoggerConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type JWTConfig struct {
	User struct {
		AccessToken struct {
			SecretKey     string        `mapstructure:"secret_key"`
			TokenLifetime time.Duration `mapstructure:"token_lifetime"`
		} `mapstructure:"access_token"`
		RefreshToken struct {
			SecretKey     string        `mapstructure:"secret_key"`
			EncryptionKey string        `mapstructure:"encryption_key"`
			TokenLifetime time.Duration `mapstructure:"token_lifetime"`
		} `mapstructure:"refresh_token"`
	} `mapstructure:"user"`
	Service struct {
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"service"`
}

type SwaggerConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	URL     string `mapstructure:"url"`
	Port    string `mapstructure:"port"`
}

type ServicesConfig struct {
	SSO struct {
		Url string `mapstructure:"addresses"`
	} `mapstructure:"sso-svc"`
	ElectorCab struct {
		Url string `mapstructure:"addresses"`
	}
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Swagger  SwaggerConfig  `mapstructure:"swagger"`
	Services ServicesConfig `mapstructure:"services"`
}

func LoadConfig() (Config, error) {
	configPath := os.Getenv("KV_VIPER_FILE")
	if configPath == "" {
		return Config{}, errors.New("KV_VIPER_FILE env var is not set")
	}
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, errors.Errorf("error reading config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, errors.Errorf("error unmarshalling config: %s", err)
	}

	return config, nil
}

func (c *Config) SetupLogger() *logrus.Logger {
	logger := logrus.New()

	lvl, err := logrus.ParseLevel(strings.ToLower(c.Logger.Level))
	if err != nil {
		logger.Warnf("invalid log level '%s', defaulting to 'info'", c.Logger.Level)
		lvl = logrus.InfoLevel
	}
	logger.SetLevel(lvl)

	switch strings.ToLower(c.Logger.Format) {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		fallthrough
	default:
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return logger
}
