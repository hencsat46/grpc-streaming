package main

import (
	"grpc-streaming/internal/controller"
	"grpc-streaming/internal/handler"
	"grpc-streaming/internal/repository"
	"log"
)

func main() {
	repo := repository.NewRepository()
	usecase := controller.NewUsecase(repo)
	handler := handler.NewHandler(usecase)
	log.Println("Starting server on port 3000")
	if err := handler.Run(":3000"); err != nil {
		log.Println(err)
	}
}
