package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Username     string
	Nickname     string
	PasswordHash string
	CreatedAt    time.Time
}

func NewUser(username, nickname, passwordHash string) *User {
	return &User{
		ID:           uuid.New(),
		Username:     username,
		Nickname:     nickname,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}
}
