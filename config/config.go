package config

import "github.com/spf13/viper"


type AppConfig struct {
	Port          string 
	MongoURI        string 
}

func SetConfig() (*AppConfig, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config := &AppConfig{
		Port:     viper.GetString("PORT"),
		MongoURI: viper.GetString("DB_URI"),
	}
	return config, nil
}