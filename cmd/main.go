package main

import (
	"log"
	"notes"
	"notes/pkg/handler"
	"notes/pkg/repository"
	"notes/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err)
	}
}
