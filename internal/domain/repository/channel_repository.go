package repository

import (
	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type ChannelRepository interface {
	Create(channel *entity.Channel) error
	GetByName(name string) (*entity.Channel, error)
	GetCreatedByUserID(userID uuid.UUID) (*[]entity.Channel, error)
	Get(limit int, offset int) (*[]entity.Channel, error)
	Delete(id uuid.UUID) error
	RemoveAllMembers(channelID uuid.UUID) error
	AddMember(userID, channelID uuid.UUID) error
	RemoveMember(userID, channelID uuid.UUID) error
	GetMembers(channelID uuid.UUID) (*[]uuid.UUID, error)
	GetMembersByUserID(userID uuid.UUID) (*[]uuid.UUID, error)
	GetByUserID(userID uuid.UUID) (*[]entity.Channel, error)
	GetMemberCounts() (map[uuid.UUID]int, error)
}
