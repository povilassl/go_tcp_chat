package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id         uuid.UUID  `db:"id"`
	UserFromID uuid.UUID  `db:"user_from_id"`
	UserToID   *uuid.UUID `db:"user_to_id"`
	ChannelID  *uuid.UUID `db:"channel_id"`
	Content    string     `db:"content"`
	CreatedAt  time.Time  `db:"created_at"`
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
