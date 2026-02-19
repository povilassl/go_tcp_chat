package repository

import (
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	GetByUsername(username string) (*entity.User, error)
}
