package repository

import usecase "grpc-streaming/internal/controller"

type repository struct {
	data map[string]int32
}

func NewRepository() usecase.RepositoryInterfaces {
	products := make(map[string]int32)
	products["Курица"] = 55
	products["Рис"] = 99
	products["Картофель"] = 11
	products["Яблоко"] = 32
	products["Груша"] = 90
	products["Банан"] = 1
	products["Овсяные хлопья"] = 100
	products["Вода"] = 781
	return &repository{data: products}
}

func (r *repository) Read() map[string]int32 {
	return r.data
}
