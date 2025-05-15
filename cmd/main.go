package main

import (
	"github.com/Danil-Ivonin/Go"
	"github.com/Danil-Ivonin/Go/pkg/handler"
	"github.com/Danil-Ivonin/Go/pkg/repository"
	"github.com/Danil-Ivonin/Go/pkg/service"
	"log"
)

func main() {
	srv := new(todo.Server)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occired while running server: %v", err)
	}

}
