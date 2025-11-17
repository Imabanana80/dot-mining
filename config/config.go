package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Fabric Loader `toml:fabric`	
	}
	Loader struct {
		Mods map[string]string `toml:mods`
	}
)

func LoadConfig() (Config) {
	file, _ := os.ReadFile("config.toml")
	var config Config 
	toml.Decode(string(file), &config)
	return config
}
