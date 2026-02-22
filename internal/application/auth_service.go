package application

import (
	"fmt"

	"github.com/povilassl/tcp_chat/internal/domain/entity"
	"github.com/povilassl/tcp_chat/internal/domain/repository"
	"github.com/povilassl/tcp_chat/internal/helpers"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users repository.UserRepository
}

func NewAuthService(users repository.UserRepository) *AuthService {
	return &AuthService{
		users: users,
	}
}

func (a *AuthService) Register(
	username string,
	password string,
	nickname *string) error {

	usernameValid, usernameMessage := helpers.IsUsernameValid(username)
	if !usernameValid {
		return fmt.Errorf("%s", usernameMessage)
	}

	existingUser, err := a.users.GetByUsername(username)
	if err == nil && existingUser != nil {
		return fmt.Errorf("Username '%s' is already taken", username)
	}

	if err != nil && err.Error() != "sql: no rows in result set" {
		return fmt.Errorf("%s", err.Error())
	}

	passwordValid, passwordMessage := helpers.IsPasswordValid(password)
	if !passwordValid {
		return fmt.Errorf("%s", passwordMessage)
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	if nickname != nil {
		nicknameValid, nicknameMessage := helpers.IsNicknameValid(*nickname)
		if !nicknameValid {
			return fmt.Errorf("%s", nicknameMessage)
		}
	} else {
		nickname = &username
	}

	user := entity.NewUser(username, *nickname, string(passwordBytes))

	err = a.users.Create(user)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return nil
}

func (a *AuthService) Login(username, password string) (*entity.User, error) {
	if username == "" || password == "" {
		return nil, fmt.Errorf("Username and password cannot be empty")
	}

	user, err := a.users.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	passwordValid := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) == nil
	if !passwordValid {
		return nil, fmt.Errorf("Invalid username or password")
	}

	return user, nil
}
