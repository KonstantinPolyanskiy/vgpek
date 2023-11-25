package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

// Инициализирует .env файл
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env не найден")
	}
}

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("ошибка чтения конфига - %s\n", err)
	}

}

// Инициализирует конфиг сервиса
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("dev")

	return viper.ReadInConfig()
}
