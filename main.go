package main

import (
	"potassium.sh/dot-mining/config"
	"potassium.sh/dot-mining/minecraft"
)



func main() {
	config := config.LoadConfig()
	minecraft.WriteOptions(config.Options, "")
}
