package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/video-validator/app/modes"
	"github.com/video-validator/config"
)

func main() {
	config.InitConfig()

	appType := viper.GetString("APP_TYPE")

	switch appType {
	case "server":
		modes.StartServer()
	default:
		log.Fatal("APP_TYPE is not set")
	}

}
