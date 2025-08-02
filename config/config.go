package config

import (
	"fmt"

	"github.com/spf13/viper"
)


type AppConfig struct {
	Port          string 
	MongoURI        string 
}

func SetConfig() (*AppConfig, error) {
	viper.SetConfigName(".env") // no extension
	viper.SetConfigType("env")  // treat it like key=value
	viper.AddConfigPath(".")    // current folder
	viper.AddConfigPath("..")   // fallback (e.g. if launched from /cmd)

	viper.AutomaticEnv() // fallback to system env vars if .env fails

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read .env file: %w", err)
	}

	config := &AppConfig{
		Port:     viper.GetString("PORT"),
		MongoURI: viper.GetString("DB_URI"),
	}
	return config, nil
}
