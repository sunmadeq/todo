package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"todo/internal/app"
	"todo/internal/handlers"
	"todo/internal/repository"
	"todo/internal/repository/database"
	"todo/internal/services"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing configuration file, %s", err.Error())
		return
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("Error loading environment file, %s", err.Error())
		return
	}

	db, err := database.NewPostgresDatabase(
		database.Config{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("database.user"),
			Database: viper.GetString("database.database"),
			Mode:     viper.GetString("database.mode"),
			Password: os.Getenv("DATABASE_PASSWORD"),
		},
	)

	if err != nil {
		logrus.Fatalf("Error initializing database connection, %s", err.Error())
		return
	}

	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	server := new(app.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("An error occurred while running the server: %s", err.Error())
		}
	}()

	logrus.Print("The application has started.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Print("The application is shutting down.")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("An error occurred while shutting down the server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("An error occurred while closing the database connection: %s", err.Error())
	}

	logrus.Print("The application has shut down.")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
