package entity

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	ID          uuid.UUID
	ChannelName string
	CreatedByID uuid.UUID
	// Members     map[uint64]*hub.Client
	CreatedAt time.Time
}

func NewChannel(name string, createdByID uuid.UUID) *Channel {
	return &Channel{
		ID:          uuid.New(),
		ChannelName: name,
		CreatedByID: createdByID,
		CreatedAt:   time.Now(),
	}
}
