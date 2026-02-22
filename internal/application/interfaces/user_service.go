package interfaces

import (
	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type UserService interface {
	Rename(user *entity.User, nickname *string) error
	GetByID(id uuid.UUID) (*entity.User, error)
}
