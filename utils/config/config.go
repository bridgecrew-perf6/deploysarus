package config

import (
	"github.com/spf13/viper"
)

var v = viper.New()

func init() {
	v.SetDefault("name", "server-1")
	v.SetDefault("platform", "github")
	// setting
	v.SetDefault("server.address", "localhost")
	v.SetDefault("server.port", 1515)
	v.SetDefault("server.path", "/webhooks")

}

// Load with custom path
func Load(path string) error {
	v.SetConfigFile(path)
	return v.ReadInConfig()
}

// Set config value
func Set(key string, value interface{}) {
	v.Set(key, value)
}

// Get gets value
func Get(key string) interface{} {
	return v.Get(key)
}

// GetInt gets int value
func GetInt(key string) int {
	return v.GetInt(key)
}

// GetString gets string value
func GetString(key string) string {
	return v.GetString(key)
}

// GetStringSlice gets strings value
func GetStringSlice(key string) []string {
	return v.GetStringSlice(key)
}

// GetStringMap gets string map value
func GetStringMap(key string) map[string]interface{} {
	return v.GetStringMap(key)
}

// UnmarshalKey is wrapper of viper Unmarshal key
func UnmarshalKey(key string, pointer interface{}) error {
	return v.UnmarshalKey(key, pointer)
}
