package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	AuthConfig AuthConfig `yaml:"auth_config"`
	DbPath     string     `yaml:"db_path"`
}

type AuthConfig struct {
	HttpPort    int    `yaml:"http_port"`
	IdentityKey string `yaml:"identity_key"`
	SigningKey  []byte `yaml:"signing_key"`
}

func (c *Config) setDefaults() error {
	if c.DbPath == "" {
		c.DbPath = "./database/sqlite.db"
	}

	return c.AuthConfig.setDefaults()
}

func (c *AuthConfig) setDefaults() error {
	if c.HttpPort == 0 {
		c.HttpPort = 8080
	} else if c.HttpPort < 1024 || c.HttpPort > 65535 {
		return errors.New("invalid http port")
	}

	if c.IdentityKey == "" {
		c.IdentityKey = "identity"
	}

	if c.SigningKey == nil {
		c.SigningKey = []byte("secret")
	}

	return nil
}

func NewConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	if err := config.setDefaults(); err != nil {
		return Config{}, err
	}

	return config, nil
}
