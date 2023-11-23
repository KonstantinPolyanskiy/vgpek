package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"practise-task-service/pkg/handler"
	"practise-task-service/pkg/storage"
	"practise-task-service/pkg/storage/postgres"
	"practise-task-service/pkg/usecase"
	"syscall"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env не найден")
	}
}

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("Ошибка в чтении cfg - %s\n", err)
	}

	savedPath := viper.GetString("folder.save")
	deletedPath := viper.GetString("folder.delete")

	err = os.Mkdir(savedPath, 0755)
	if err != nil {
		log.Printf("Ошибка в создании папки - %s\n", err)
	}

	err = os.Mkdir(deletedPath, 0755)
	if err != nil {
		log.Printf("Ошибка в создании папки - %s\n", err)
	}

	db, err := postgres.New(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   "postgres",
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Ошибка в инициалзиации бд - %s", err)
	}
	repo := storage.New(db, savedPath, deletedPath)
	service := usecase.New(repo)
	handlers := handler.New(service)

	srv := &http.Server{
		Addr:         viper.GetString("http_server.address"),
		Handler:      handlers.Init(),
		ReadTimeout:  viper.GetDuration("http_server.timeout"),
		WriteTimeout: viper.GetDuration("http_server.timeout"),
		IdleTimeout:  viper.GetDuration("http_server.idle_timeout"),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Запуск сервера")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("ошибка в запускее сервера - %s\n", err)
		}
	}()
	<-done
	log.Println("Остановка сервера")

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("dev")

	return viper.ReadInConfig()
}
