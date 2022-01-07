package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Set method
// you can set new key in switch for manage config with config server
func (g *Config) Set(key string, query []byte) error {
	if err := json.Unmarshal(query, &Confs); err != nil {
		return err
	}
	return nil
}

func (g *Config) GetDebug() bool {
	return g.Debug
}

func (g *Config) SetDebug(debug bool) {
	g.Debug = debug
}

// Load returns configs
func (g *Config) Load(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return g.file(path)
	}

	return fmt.Errorf("file not exists")
}

// file func
func (g *Config) file(path string) error {

	// name of config file (without extension)
	// REQUIRED if the config file does not have the extension in the name
	// path to look for the config file in
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return err

	}

	return viper.Unmarshal(&Confs)
}
