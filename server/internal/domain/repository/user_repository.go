package repository

import (
	"context"

	entities "github.com/povilassl/tcp_chat/server/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User)
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
}
