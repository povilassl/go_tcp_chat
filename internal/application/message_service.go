package application

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
	"github.com/povilassl/tcp_chat/internal/domain/repository"
)

type MessageService struct {
	messages repository.MessageRepository
	channels repository.ChannelRepository
}

func NewMessageService(
	messages repository.MessageRepository,
	channels repository.ChannelRepository) *MessageService {
	return &MessageService{
		messages: messages,
		channels: channels,
	}
}

func (a *MessageService) Create(
	userFromID uuid.UUID,
	userToID *uuid.UUID,
	channelName *string,
	content string) (*entity.Message, error) {
	if userFromID == uuid.Nil {
		return nil, fmt.Errorf("Message must have UserFromID set")
	}

	if (userToID != nil && channelName != nil) ||
		(userToID == nil && channelName == nil) {
		return nil, fmt.Errorf("Message must have either UserToID or ChannelName set")
	}

	var channelID *uuid.UUID
	if channelName != nil {
		channel, err := a.channels.GetByName(*channelName)
		if err != nil {
			return nil, fmt.Errorf("%s", err.Error())
		}

		channelID = &channel.ID
	}

	message := entity.NewMessage(content, userFromID, userToID, channelID)
	err := a.messages.Create(&message)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return &message, nil
}

//TODO delete?
