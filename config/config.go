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
		DBHost:        viper.GetString("DB_HOST"),
		DBPort:        viper.GetInt("DB_PORT"),
		DBUser:        viper.GetString("DB_USER"),
		DBPassword:    viper.GetString("DB_PASSWORD"),
		DBName:        viper.GetString("DB_NAME"),
		DBSSLMode:     viper.GetString("DB_SSL"),
		ServerPort:    viper.GetString("SERVER_PORT"),
		SecretKey:     viper.GetString("SECRET_KEY"),
		Domain:        viper.GetString(constants.Domain),
		MigrationsDir: viper.GetString("MIGRATION_DIR"),
	}

	log.Printf("Config Loaded: %+v\n", AppConfig)
}
