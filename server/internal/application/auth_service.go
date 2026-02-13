package application

import (
	"github.com/povilassl/tcp_chat/server/internal/domain/entity"
	"github.com/povilassl/tcp_chat/server/internal/domain/repository"
)

type AuthService struct {
	users repository.UserRepository
}

func (AuthService) Register(username, password string) error

func (AuthService) Login(username, password string) (entity.User, error)
