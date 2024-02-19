package main

import (
	"context"
	"log"
	"time"

	"github.com/hencsat46/protos-streaming/gen/go/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":3000", opts...)
	if err != nil {
		log.Println(err)
		return
	}

	client := store.NewStoreClient(conn)

	updateStream, err := client.UpdateProducts(context.Background())

	if err != nil {
		log.Println(err)
		return
	}

	products := make(map[string]int32)
	products["Булочка с сосискою"] = 1
	products["Чай бутерброд"] = 1
	products["И пряники то чаю"] = 1

	for key, value := range products {
		time.Sleep(5 * time.Second)
		if err := updateStream.Send(&store.UpdateRequest{Name: key, Count: value}); err != nil {
			log.Println(err)
			return
		}
	}

	updateResult, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(updateResult)
}
