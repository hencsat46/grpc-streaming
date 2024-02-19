package handler

import (
	"log"
	"net"

	store "github.com/hencsat46/protos-streaming/gen/go/store"
	"google.golang.org/grpc"
)

type handler struct {
	products map[string]int32
	store.UnimplementedStoreServer
}

func NewHandler() *handler {
	products := make(map[string]int32)
	products["Курица"] = 55
	products["Рис"] = 99
	products["Картофель"] = 11
	products["Яблоко"] = 32
	products["Груша"] = 90
	products["Банан"] = 1
	products["Овсяные хлопья"] = 100
	products["Вода"] = 781
	return &handler{products: products}
}

func (h *handler) Run(port string) error {
	listener, err := net.Listen("tcp4", port)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	h.register(server)

	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (h *handler) register(gRPC *grpc.Server) {
	store.RegisterStoreServer(gRPC, &handler{products: h.products})
}

func (h *handler) GetProducts(response *store.OrderRequest, stream store.Store_GetProductsServer) error {
	for key, value := range h.products {
		if err := stream.Send(&store.Order{Name: key, Count: value}); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
