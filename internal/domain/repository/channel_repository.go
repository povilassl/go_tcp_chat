package repository

import (
	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

// TODO: leave / join?
type ChannelRepository interface {
	Create(channel *entity.Channel) error
	GetByName(name string) (*entity.Channel, error)
	GetByUserID(userID uuid.UUID) (*[]entity.Channel, error)
	Delete(id uuid.UUID) error
}
