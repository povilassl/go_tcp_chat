package entity

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	ID          uuid.UUID `db:"id"`
	ChannelName string    `db:"channel_name"`
	CreatedByID uuid.UUID `db:"created_by_id"`
	CreatedAt   time.Time `db:"created_at"`
}

func NewChannel(name string, createdByID uuid.UUID) *Channel {
	return &Channel{
		ID:          uuid.New(),
		ChannelName: name,
		CreatedByID: createdByID,
		CreatedAt:   time.Now(),
	}
}
