package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	// Example Variable
	APIPort   string
	RedisHost string
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("lbc")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("lbc")
	v.SetDefault("APIPort", "4242")
	v.SetDefault("RedisHost", "127.0.0.1:6379")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
