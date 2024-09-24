package main

import (
	"log/slog"

	"github.com/spf13/viper"
	"github.com/video-validator/cmd/modes"
	"github.com/video-validator/config"
	"github.com/video-validator/internal/helper"
	"github.com/video-validator/pkg"
)

func main() {
	config.InitConfig()
	pkg.LoggerInit()
	pkg.DbInit()
	helper.HttpInit()

	appType := viper.GetString("APP_TYPE")

	switch appType {
	case "server":
		modes.StartServer()
	case "singular":
		modes.StartSingular()
	case "consumer":
		// modes.StartConsumer()
	default:
		slog.Error("APP_TYPE is not set")
	}
}
