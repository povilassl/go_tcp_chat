package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	Nickname     string    `db:"nickname"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
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
