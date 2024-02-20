package main

import (
	"context"
	"errors"
	"io"
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

	Send(client)
}

func Send(client store.StoreClient) {
	products := make(map[string]int32)
	products["Булочка с сосискою"] = 1
	products["Чай бутерброд"] = 1
	products["И пряники то чаю"] = 1

	stream, err := client.ThrowProducts(context.Background())

	if err != nil {
		log.Println(err)
		return
	}

	for key, value := range products {
		time.Sleep(5 * time.Second)
		if err := stream.Send(&store.UpdateRequest{Name: key, Count: value}); err != nil {
			log.Println(err)
			return
		}
	}

	stream.CloseSend()

	for {
		result, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		log.Println(result)
	}

}
