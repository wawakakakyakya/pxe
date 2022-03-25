package config

import (
	"flag"
)

var configPath string

func setArgs() {
	flag.StringVar(&configPath, "src", "", "default ./")
	flag.Parse()
}
