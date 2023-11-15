package main

import (
	"log"
	"mod/internal/handler"
	"mod/internal/storage"
	"mod/internal/storage/postgres"
	"mod/internal/usecase"
)

func main() {
	db, err := postgres.New(postgres.Config{
		Host:     "",
		Port:     "",
		Username: "",
		Password: "",
		DBName:   "",
		SSLMode:  "",
	})
	if err != nil {
		log.Fatalf("Ошибка в инициалзиации бд - %s", err)
	}
	repository := storage.New(db)
	service := usecase.New(repository)
	handler := handler.New(service)
}
