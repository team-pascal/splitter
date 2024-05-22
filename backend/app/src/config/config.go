// For using environment variable

package config

import "github.com/spf13/viper"

type Config struct {
	// For connecting DB
	DB struct {
		User     string
		Password string
		DBName   string
		Host     string
		Port     string
	}
}

func LoadConfig() Config {

	// Import .env file
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var config Config

	// Create structure for using environment variable
	config.DB.User = viper.GetString("POSTGRES_USER")
	config.DB.Password = viper.GetString("POSTGRES_PASSWORD")
	config.DB.DBName = viper.GetString("POSTGRES_DB")
	config.DB.Host = viper.GetString("POSTGRES_HOST")
	config.DB.Port = viper.GetString("POSTGRES_PORT")

	return config
}
