package mysql

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *entity.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, username, nickname, password_hash, created_at) VALUES (?, ?, ?, ?, ?)",
		user.ID,
		user.Username,
		user.Nickname,
		user.PasswordHash,
		user.CreatedAt,
	)

	return err
}

func (r *UserRepository) GetByUsername(username string) (*entity.User, error) {
	user := entity.User{}
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = ?", username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByID(id uuid.UUID) (*entity.User, error) {
	user := entity.User{}
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	_, err := r.db.Exec("UPDATE users SET username = ?, nickname = ?, password_hash = ? WHERE id = ?",
		user.Username,
		user.Nickname,
		user.PasswordHash,
		user.ID)

	return err
}
