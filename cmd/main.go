package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occured while initializating config: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error while loading env variables: %s", err)
	}

	db, err := repository.NewDBConnection(os.Getenv("DB_CONNECT"))
	if err != nil {
		log.Fatalf("error occured while initializating database: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
