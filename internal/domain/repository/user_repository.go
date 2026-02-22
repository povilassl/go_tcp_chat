package repository

import (
	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	GetByID(id uuid.UUID) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
}
