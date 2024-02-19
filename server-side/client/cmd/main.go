package main

import (
	"context"
	"errors"
	"io"
	"log"

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

	searchStream, _ := client.GetProducts(context.Background(), &store.OrderRequest{})

	for {
		searchOrder, err := searchStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		log.Println(searchOrder)
	}
}
