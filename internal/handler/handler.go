package handler

import (
	"log"
	"net"
	"time"

	store "github.com/hencsat46/protos-streaming/gen/go/store"
	"google.golang.org/grpc"
)

type handler struct {
	usecase UsecaseInterfaces
	store.UnimplementedStoreServer
}

type UsecaseInterfaces interface {
	Read() map[string]int32
}

func NewHandler(usecase UsecaseInterfaces) *handler {

	return &handler{usecase: usecase}
}

func (h *handler) Run(port string) error {
	listener, err := net.Listen("tcp4", port)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	h.register(server, h.usecase)

	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (h *handler) register(gRPC *grpc.Server, usecase UsecaseInterfaces) {
	store.RegisterStoreServer(gRPC, &handler{usecase: usecase})
}

func (h *handler) GetProducts(response *store.OrderRequest, stream store.Store_GetProductsServer) error {

	products := h.usecase.Read()

	for key, value := range products {
		time.Sleep(5 * time.Second)
		if err := stream.Send(&store.Order{Name: key, Count: value}); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
