package service

import (
	"crypto/sha256"
	"fmt"
	todo "github.com/Danil-Ivonin/Go"
	"github.com/Danil-Ivonin/Go/pkg/repository"
)

const salt = "1kl2j3h4"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
