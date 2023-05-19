package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error occured while initializating config: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error while loading env variables: %s", err)
	}

	db, err := repository.NewDBConnection(os.Getenv("DB_CONNECT"))
	if err != nil {
		logrus.Fatalf("error occured while initializating database: %s", err)
	}
	defer db.Close()

	repos := repository.NewRepository(db.Client, os.Getenv("DB_NAME"))
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
