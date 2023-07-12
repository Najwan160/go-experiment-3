package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Env Environment
)

type Environment struct {
	Port            int   `mapstructure:"PORT"`
	DevelopmentMode bool  `mapstructure:"DEVELOPMENT_MODE"`
	DB              EnvDB `mapstructure:",squash"`
	JWT             JWT   `mapstructure:",squash"`
}

type EnvDB struct {
	Host            string `mapstructure:"DB_HOST"`
	Port            int    `mapstructure:"DB_PORT"`
	Username        string `mapstructure:"DB_USERNAME"`
	Password        string `mapstructure:"DB_PASSWORD"`
	Schema          string `mapstructure:"DB_SCHEMA"`
	Debug           bool   `mapstructure:"DB_DEBUG"`
	CreateBatchSize int    `mapstructure:"DB_CREATE_BATCH_SIZE"`
}

type JWT struct {
	TTL          int    `mapstructure:"TOKEN_TTL"`
	SignatureKey string `mapstructure:"TOKEN_SIGNATURE_KEY"`
}

func LoadEnv(path string) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("unable to read env: %v", err))
	}

	if err := viper.Unmarshal(&Env); err != nil {
		panic(fmt.Sprintf("unable to parse env: %v", err))
	}
}
