package main

import (
	"potassium.sh/dot-mining/config"
	"potassium.sh/dot-mining/modrinth"
)



func main() {
	config := config.LoadConfig()
	modrinth.LoadMods(config.Mods)
}
