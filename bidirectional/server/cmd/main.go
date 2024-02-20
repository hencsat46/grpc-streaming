package main

import (
	"grpc-streaming/bidirectional/server/internal/controller"
	"grpc-streaming/bidirectional/server/internal/handler"
	"grpc-streaming/bidirectional/server/internal/repository"
	"log"
)

func main() {
	repo := repository.NewRepository()
	usecase := controller.NewUsecase(repo)
	handler := handler.NewHandler(usecase)
	log.Println("Server started on port 3000")

	if err := handler.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
