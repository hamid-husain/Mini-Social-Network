package config

import (
	"mini-social-network/constants"

	"github.com/spf13/viper"

	"log"
)

type Config struct {
	DBHost        string
	DBPort        int
	DBUser        string
	DBPassword    string
	DBName        string
	DBSSLMode     string
	ServerPort    string
	SecretKey     string
	MigrationsDir string
	Domain        string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = Config{
		DBHost:        viper.GetString(constants.DBHost),
		DBPort:        viper.GetInt(constants.DBPort),
		DBUser:        viper.GetString(constants.DBUser),
		DBPassword:    viper.GetString(constants.DBPassword),
		DBName:        viper.GetString(constants.DBName),
		DBSSLMode:     viper.GetString(constants.DBSSLMode),
		ServerPort:    viper.GetString(constants.ServerPort),
		SecretKey:     viper.GetString(constants.SecretKey),
		Domain:        viper.GetString(constants.Domain),
		MigrationsDir: viper.GetString(constants.MigrationDir),
	}

	log.Printf("Config Loaded: %+v\n", AppConfig)
}
