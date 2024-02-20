package controller

import "grpc-streaming/bidirectional/server/internal/handler"

type usecase struct {
	repo RepositroyInterfaces
}

type RepositroyInterfaces interface {
	Add(string, int32)
	Get() map[string]int32
}

func NewUsecase(repo RepositroyInterfaces) handler.UsecaseInterfaces {
	return &usecase{repo: repo}
}

func (u *usecase) Add(name string, count int32) {
	u.repo.Add(name, count)
}

func (u *usecase) Get() map[string]int32 {
	return u.repo.Get()
}
