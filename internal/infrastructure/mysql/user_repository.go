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
	_, err := r.db.Exec("INSERT INTO users (id, username, nickname, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)",
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
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1", username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	_, err := r.db.Exec("UPDATE users SET username = $1, nickname = $2, password_hash = $3 WHERE id = $4",
		user.Username,
		user.Nickname,
		user.PasswordHash,
		user.ID)

	return err
}
