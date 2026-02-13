package entity

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	ID   uuid.UUID
	Name string
	// Members   map[uint64]*Client //TODO?
	CreatedBy *User
	CreatedAt time.Time
}
