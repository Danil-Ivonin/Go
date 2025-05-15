package main

import (
	todo "github.com/Danil-Ivonin/Go"
	"github.com/Danil-Ivonin/Go/pkg/handler"
	"log"
)

func main() {
	srv := new(todo.Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occired while running server: %v", err)
	}

}
