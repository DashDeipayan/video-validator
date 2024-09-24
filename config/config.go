package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		fmt.Println("APP_ENV is not set, defaulting to 'local'")
		appEnv = "local"
		fmt.Printf("APP_ENV is set to here: %s\n", appEnv)
	}
	viper.SetConfigName(appEnv)

	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	config := make(map[string]interface{})
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
	}

	processKeys(config, "")
	viper.AutomaticEnv()
}

func processKeys(config map[string]interface{}, prefix string) {
	for key, value := range config {
		upperKey := strings.ToUpper(strings.ReplaceAll(prefix+key, ".", "_"))
		switch v := value.(type) {
		case map[string]interface{}:
			processKeys(v, upperKey+"_")
		default:
			viper.Set(upperKey, v)
			if os.Getenv(upperKey) == "" {
				os.Setenv(upperKey, fmt.Sprintf("%v", v))
			}
		}
	}
}
