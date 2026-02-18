package application

import (
	"fmt"

	"github.com/povilassl/tcp_chat/internal/domain/entity"
	"github.com/povilassl/tcp_chat/internal/domain/repository"
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

// TODO add service interfaces
func (a *AuthService) Register(username, password string) error {

	usernameValid, usernameMessage := isUsernameValid(username)
	if !usernameValid {
		return fmt.Errorf("Error registering user: %s", usernameMessage)
	}

	passwordValid, passwordMessage := isPasswordValid(password)
	if !passwordValid {
		return fmt.Errorf("Error registering user: %s", passwordMessage)
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return fmt.Errorf("Error registering user: %s", err.Error())
	}

	user := entity.NewUser(username, string(passwordBytes))

	err = a.users.Create(user)
	if err != nil {
		return fmt.Errorf("Error registering user: %s", err.Error())
	}

	return nil
}

func (a *AuthService) Login(username, password string) (*entity.User, error) {

	if username == "" || password == "" {
		return nil, fmt.Errorf("Error logging in: username and password cannot be empty")
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, fmt.Errorf("Error logging in, could not hash password: %s", err.Error())
	}

	user, err := a.users.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("Error logging in: %s", err.Error())
	}

	passwordValid := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(passwordBytes)) == nil
	if !passwordValid {
		return nil, fmt.Errorf("Error logging in: invalid username or password")
	}

	return user, nil
}
