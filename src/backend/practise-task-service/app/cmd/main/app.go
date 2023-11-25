package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"practise-task-service/internal/handler"
	"practise-task-service/internal/storage"
	"practise-task-service/internal/storage/postgres"
	"practise-task-service/internal/usecase"
	log_err "practise-task-service/pkg/logger/error"
	slog_dev "practise-task-service/pkg/logger/handlers/slog-dev"
	"syscall"
)

const (
	envDev  = "dev"
	envProd = "prod"
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
		log.Fatalf("Ошибка в чтении cfg - %s\n", err)
	}

	savedPath := viper.GetString("folder.save")
	deletedPath := viper.GetString("folder.delete")

	logger := setupLogger(viper.GetString("env"))

	logger.Info(
		"запуск practice-task microservice",
		slog.String("окружение", viper.GetString("env")))

	err = os.Mkdir(savedPath, 0755)
	if err != nil {
		logger.Warn("ошибка в создании папки", log_err.Err(err))
	}

	err = os.Mkdir(deletedPath, 0755)
	if err != nil {
		logger.Warn("ошибка в создании папки", log_err.Err(err))
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
		logger.Error("ошибка в инициалзиации базы данных", log_err.Err(err))
		panic("error init db")
	}
	repo := storage.New(db, savedPath, deletedPath, logger)
	service := usecase.New(repo)
	handlers := handler.New(service, logger)

	srv := &http.Server{
		Addr:         viper.GetString("http_server.address"),
		Handler:      handlers.Init(),
		ReadTimeout:  viper.GetDuration("http_server.timeout"),
		WriteTimeout: viper.GetDuration("http_server.timeout"),
		IdleTimeout:  viper.GetDuration("http_server.idle_timeout"),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("запуск сервера на порту", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("ошибка в запуске сервера", log_err.Err(err))
		}
	}()
	<-done
	logger.Info("остановка сервера")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("dev")

	return viper.ReadInConfig()
}

func setupLogger(env string) *slog.Logger {
	//TODO: настраивать логер в зависимости от окружения
	return slog_dev.SetupDevSlog()
}
