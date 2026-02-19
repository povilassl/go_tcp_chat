package application

import (
	"fmt"

	"github.com/povilassl/tcp_chat/internal/domain/entity"
	"github.com/povilassl/tcp_chat/internal/domain/repository"
)

type ChannelService struct {
	channels repository.ChannelRepository
	messages repository.MessageRepository
}

func NewChannelService(
	channels repository.ChannelRepository,
	messages repository.MessageRepository) *ChannelService {
	return &ChannelService{
		channels: channels,
		messages: messages,
	}
}

func (a *ChannelService) Create(name string, user *entity.User) error {
	if user == nil {
		return fmt.Errorf("You must be logged in to create a channel")
	}

	nameValid, nameMessage := isChannelNameValid(name)
	if !nameValid {
		return fmt.Errorf("%s", nameMessage)
	}

	channelExists, err := a.channels.GetByName(name)

	//TODO double check err message
	if err != nil && err.Error() != "sql: no rows in result set" {
		return fmt.Errorf("%s", err.Error())
	}

	if channelExists != nil {
		return fmt.Errorf("Channel with this name already exists")
	}

	channelsByUser, err := a.channels.GetByUserID(user.ID)
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

	err = a.messages.DeleteByChannelID(channel.ID)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return a.channels.Delete(channel.ID)
}

func (a *ChannelService) Get(limit int, offset int) (*[]entity.Channel, error) {
	if limit <= 0 {
		limit = 100
	}

	if offset < 0 {
		offset = 0
	}

	return a.channels.Get(limit, offset)
}
