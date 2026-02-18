package mysql

import (
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
	_, err := r.db.Exec("INSERT INTO users (id, username, password_hash, created_at) VALUES ($1, $2, $3, $4)",
		user.ID,
		user.Username,
		user.PasswordHash,
		user.CreatedAt,
	)

	return err
}

func (r *UserRepository) GetByUsername(username string) (*entity.User, error) {
	user := entity.User{}
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1", username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
