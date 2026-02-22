package application

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
	"github.com/povilassl/tcp_chat/internal/domain/repository"
	"github.com/povilassl/tcp_chat/internal/helpers"
)

type UserService struct {
	users repository.UserRepository
}

func NewUserService(users repository.UserRepository) *UserService {
	return &UserService{
		users: users,
	}
}

func (a *UserService) Rename(user *entity.User, nickname *string) error {
	nicknameValid, nicknameMessage := helpers.IsNicknameValid(*nickname)
	if !nicknameValid {
		return fmt.Errorf("%s", nicknameMessage)
	}

	user.Nickname = *nickname

	err := a.users.Update(user)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return nil
}

func (a *UserService) GetByID(id uuid.UUID) (*entity.User, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("Invalid user ID")
	}

	return a.users.GetByID(id)
}
