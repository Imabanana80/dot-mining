package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"potassium.sh/dot-mining/minecraft"
)

type (
	Config struct {
		Options minecraft.Options `toml:options`
		Fabric Loader `toml:fabric`	
		Mods map[string]string `toml:mods`
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
