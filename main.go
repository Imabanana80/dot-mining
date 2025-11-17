package main

import (
	"fmt"
	
	"potassium.sh/dot-mining/config"
)



func main() {
	config := config.LoadConfig()
	for modName, modID := range config.Fabric.Mods {
		fmt.Printf("- %s: %s\n", modName, modID)
	}
}
