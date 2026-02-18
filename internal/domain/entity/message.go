package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id         uuid.UUID
	ChannelID  uuid.UUID
	UserFromID uuid.UUID
	UserToID   uuid.UUID
	Content    string
	CreatedAt  time.Time
}
