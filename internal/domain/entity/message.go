package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id         uuid.UUID
	UserFromID uuid.UUID
	UserToID   *uuid.UUID
	ChannelID  *uuid.UUID
	Content    string
	CreatedAt  time.Time
}

func NewMessage(
	content string,
	userFromID uuid.UUID,
	channelID *uuid.UUID,
	userToID *uuid.UUID) Message {
	return Message{
		Id:         uuid.New(),
		Content:    content,
		UserFromID: userFromID,
		ChannelID:  channelID,
		UserToID:   userToID,
		CreatedAt:  time.Now(),
	}
}
