package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Password  string
	CreatedAt time.Time
}
