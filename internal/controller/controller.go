package controller

import "grpc-streaming/internal/handler"

type usecase struct {
	repo RepositoryInterfaces
}

type RepositoryInterfaces interface {
	Read() map[string]int32
}

func NewUsecase(repo RepositoryInterfaces) handler.UsecaseInterfaces {
	return &usecase{repo: repo}
}

func (u *usecase) Read() map[string]int32 {
	return u.repo.Read()
}
