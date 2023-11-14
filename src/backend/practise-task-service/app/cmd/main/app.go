package main

import (
	"log"
	"mod/internal/storage/postgres"
)

func main() {
	cfg := postgres.Config{
		Host:     "",
		Port:     "",
		Username: "",
		Password: "",
		DBName:   "",
		SSLMode:  "",
	}
	_, err := postgres.New(cfg)
	if err != nil {
		log.Println("Ошибка в получении подключения к БД - ", err)
	}
}
