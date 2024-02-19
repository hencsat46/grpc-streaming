package main

import (
	"grpc-streaming/internal/handler"
	"log"
)

func main() {
	handler := handler.NewHandler()
	log.Println("Starting server on port 3000")
	if err := handler.Run(":3000"); err != nil {
		log.Println(err)
	}
}
