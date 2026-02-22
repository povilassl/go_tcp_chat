package interfaces

import "github.com/povilassl/tcp_chat/internal/domain/entity"

type AuthService interface {
	Register(username string, password string, nickname *string) error
	Login(username, password string) (*entity.User, error)
}
