package service

import (
	todo "github.com/Danil-Ivonin/Go"
	"github.com/Danil-Ivonin/Go/pkg/repository"
)

type Authorization interface {
	CreateUser(todo.User) (int, error)
}
type TodoList interface{}
type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
