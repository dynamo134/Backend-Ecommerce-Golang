package config

import "github.com/spf13/viper"


type Config struct {
	Port          string 
	MongoURI        string 
}

func SetConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config := &Config{
		Port:     viper.GetString("PORT"),
		MongoURI: viper.GetString("DB_URI"),
	}
	return config, nil
}