package main

import (
	"github.com/artemiyKew/json-rpc-lamoda/internal/app"
)

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}
