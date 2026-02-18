package repository

import (
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type MessageRepository interface {
	Create(message *entity.Message) error
}
