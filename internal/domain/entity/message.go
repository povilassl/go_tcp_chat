package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID
	ChannelID uuid.UUID
	UserID    uuid.UUID
	Content   string
	CreatedAt time.Time
}
