package config

import "github.com/spf13/viper"

type Env struct {
	AppMode          string `mapstructure:"APP_MODE"`
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	ServerPort       string `mapstructure:"SERVER_PORT"`
	BasePrefixUrl    string `mapstructure:"BASE_PREFIX_URL"`
	ContextTimeout   int    `mapstructure:"CONTEXT_TIMEOUT"`
	DatabaseDriver   string `mapstructure:"DATABASE_DRIVER"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabaseUsername string `mapstructure:"DATABASE_USERNAME"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	DatabaseSSLMode  string `mapstructure:"DATABASE_SSL_MODE"`
	DatabaseTimeZone string `mapstructure:"DATABASE_TIME_ZONE"`
	IdleTimeout      int    `mapstructure:"IDLE_TIMEOUT"`
	MaxLifetime      int    `mapstructure:"MAX_LIFETIME"`
	MaxOpenConns     int    `mapstructure:"MAX_OPEN_CONNS"`
	MaxIdleConns     int    `mapstructure:"MAX_IDLE_CONNS"`
	LogLevel         int    `mapstructure:"LOG_LEVEL"`
}

func NewEnv() (*Env, error) {
	var env *Env

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&env); err != nil {
		return nil, err
	}

	return env, nil
}
