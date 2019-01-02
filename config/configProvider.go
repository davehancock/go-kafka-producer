package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// InitConfigProvider initializes a config provider.
// This implementation uses viper for its configuration management backend.
func InitConfigProvider() Provider {

	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n ", err))
	}

	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	v.AutomaticEnv()

	return v
}

// Provider is an abstraction to allow the reading of configuration values from the backend.
type Provider interface {
	Get(key string) interface{}
	GetString(key string) string
}
