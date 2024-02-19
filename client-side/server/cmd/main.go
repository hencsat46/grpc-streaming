package main

import (
	"grpc-streaming/client-side/server/internal/handler"
	"log"
)

func main() {
	handler := handler.NewHandler()

	log.Println("Server started on port 3000")
	if err := handler.Run(":3000"); err != nil {
		log.Println(err)
		return
	}

}
