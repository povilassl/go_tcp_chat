package interfaces

import (
	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type MessageService interface {
	Create(userFromID uuid.UUID, userToID *uuid.UUID, channelName *string, content string) (*entity.Message, error)
}
