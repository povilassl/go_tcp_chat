package repository

import (
	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type MessageRepository interface {
	Create(message *entity.Message) error
	DeleteByChannelID(channelID uuid.UUID) error
}
