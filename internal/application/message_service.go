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
	channelName string,
	userFromID uuid.UUID,
	userToID uuid.UUID,
	content string) error {

	if userFromID == uuid.Nil {
		return fmt.Errorf("Message must have UserFromID set")
	}

	if userToID != uuid.Nil && channelName != "" {
		return fmt.Errorf("Message cannot have both UserToID and ChannelName set")
	}

	nameValid, nameMessage := isChannelNameValid(channelName)
	if !nameValid {
		return fmt.Errorf("%s", nameMessage)
	}

	channelExists, err := a.channels.GetByName(channelName)

	//TODO double check err message
	if err != nil && err.Error() != "sql: no rows in result set" {
		return fmt.Errorf("%s", err.Error())
	}

	if channelExists != nil {
		return fmt.Errorf("Channel with this name already exists")
	}

	channelsByUser, err := a.channels.GetByUserID(userFromID)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	if len(*channelsByUser) >= 3 {
		return fmt.Errorf("You have reached the limit of channels you can create, current maximum limit is 3")
	}

	channel := entity.NewChannel(name, user.ID)
	return a.channels.Create(channel)
}

func (a *ChannelService) Delete(name string, user *entity.User) error {
	if user == nil {
		return fmt.Errorf("You must be logged in to delete a channel")
	}

	channel, err := a.channels.GetByName(name)

	//TODO double check err message
	if err != nil && err.Error() != "sql: no rows in result set" {
		return fmt.Errorf("%s", err.Error())
	}

	if channel == nil {
		return fmt.Errorf("Channel by name '%s' does not exist", name)
	}

	if channel.CreatedByID != user.ID {
		return fmt.Errorf("You do not have permissions to delete channel '%s'", name)
	}

	return a.channels.Delete(channel.ID)
}
