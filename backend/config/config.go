package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Server struct {
	Port int `mapstructure:"port"`
}

type Database struct {
	DBHost     string `mapstructure:"host"`
	DBPort     int    `mapstructure:"port"`
	DBUser     string `mapstructure:"user"`
	DBPassword string `mapstructure:"password"`
	DBName     string `mapstructure:"dbname"`
	SSLMode    string `mapstructure:"sslmode"`
}

type Configuration struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
}

var cfg Configuration

func initConfig() {
	viper.AddConfigPath("./backend/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config: ", err)
		os.Exit(0)
	}
}

func InitConfigDsn() string {
	initConfig()

	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("[Config][InitConfigDsn] Unable to decode into str")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.SSLMode)

	fmt.Println(dsn)
	return dsn

}

func InitConfigPort() string {
	initConfig()

	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("[Conifg][InitConfigPort] Unable to decode into str")
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	fmt.Println(port)

	return port

}
