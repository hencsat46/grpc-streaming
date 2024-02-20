package repository

import "grpc-streaming/bidirectional/server/internal/controller"

type repository struct {
	data map[string]int32
}

func NewRepository() controller.RepositroyInterfaces {
	data := make(map[string]int32)
	return &repository{data: data}
}

func (r *repository) Add(name string, count int32) {
	r.data[name] = count
}

func (r *repository) Get() map[string]int32 {
	return r.data
}
