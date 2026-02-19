package application

import (
	"fmt"

	"github.com/povilassl/tcp_chat/internal/domain/entity"
	"github.com/povilassl/tcp_chat/internal/domain/repository"
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
	nicknameValid, nicknameMessage := isNicknameValid(*nickname)
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
