package yoconfig

import (
	"github.com/spf13/viper"
)

type Config struct {
	Viper *viper.Viper
}

// This function will load the environment variables for the keys preset in the config file
func Load(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigName("config.json")
	v.SetConfigType("json")
	// We need to get the env values to the config keys
	v.AutomaticEnv()

	if len(path) > 0 {
		v.AddConfigPath(path)
	}
	// Add additional config paths for working dir and /etc/yobirthday
	v.AddConfigPath(".")
	v.AddConfigPath("/etc/yobirthday/")

	// Load the config values into memory
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return nil, err
		}
	}

	config := &Config{Viper: v}
	return config, nil
}

// This function is a wrapper for vipers GetString function
// For ease of use
func (c *Config) GetString(key string) string {
	return c.Viper.GetString(key)
}

func (c *Config) Get(key string) string {
	value := c.Viper.Get(key)
	return value.(string)
}