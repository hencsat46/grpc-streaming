package handler

import (
	"errors"
	"io"
	"log"
	"net"

	store "github.com/hencsat46/protos-streaming/gen/go/store"
	"google.golang.org/grpc"
)

type handler struct {
	usecase UsecaseInterfaces
	store.UnimplementedStoreServer
}

type UsecaseInterfaces interface {
	Add(string, int32)
	Get() map[string]int32
}

func NewHandler(usecase UsecaseInterfaces) *handler {
	return &handler{usecase: usecase}
}

func (h *handler) Run(port string) error {
	listener, err := net.Listen("tcp4", port)
	if err != nil {
		log.Println(err)
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

func (h *handler) UpdateProducts(stream store.Store_UpdateProductsServer) error {
	for {
		streamOrder, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			log.Println("Streaming is over")
			log.Println(h.usecase.Get())
			return stream.SendAndClose(&store.UpdateStatus{Status: "Ok"})
		}
		log.Println(streamOrder)
		h.usecase.Add(streamOrder.GetName(), streamOrder.GetCount())
	}

}
