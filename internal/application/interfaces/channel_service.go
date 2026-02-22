package interfaces

import (
	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type ChannelService interface {
	Create(name string, user *entity.User) error
	Delete(name string, user *entity.User) error
	Get(limit int, offset int) (*[]entity.Channel, error)
	GetByName(name string) (*entity.Channel, error)
	AddMember(userID, channelID uuid.UUID) error
	RemoveMember(userID, channelID uuid.UUID) error
	GetMembers(channelID uuid.UUID) (*[]uuid.UUID, error)
	GetMembersByUserID(userID uuid.UUID) (*[]uuid.UUID, error)
	GetByUserID(userID uuid.UUID) (*[]entity.Channel, error)
	GetMemberCounts() (map[uuid.UUID]int, error)
}
