package repository

import (
	"context"

	"github.com/povilassl/tcp_chat/server/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
}
