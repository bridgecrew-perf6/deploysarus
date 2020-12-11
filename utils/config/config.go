package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var v = viper.New()

// LoadConfigFile with custome path
func LoadConfigFile(path string) {
	v.SetConfigFile(path)
	v.ReadInConfig()
}

// Set config value
func Set(key string, value interface{}) {
	v.Set(key, value)
}

// Get config value
func Get(key string) interface{} {
	fmt.Printf("%#v", v)
	return v.Get(key)
}
