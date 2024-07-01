package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	LineClientID     string
	LineClientSecret string
	NgrokURL         string
	LineRedirectURI  string
	Secret           string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	AppConfig = Config{
		LineClientID:     viper.GetString("LINE_CLIENT_ID"),
		LineClientSecret: viper.GetString("LINE_CLIENT_SECRET"),
		NgrokURL:         viper.GetString("NGROK_URL"),
		LineRedirectURI:  viper.GetString("NGROK_URL") + "/api/line/callback",
		Secret:           viper.GetString("SECRET"),
	}
}
