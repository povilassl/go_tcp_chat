package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type MessageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) Create(message *entity.Message) error {
	_, err := r.db.Exec("INSERT INTO messages (id, channel_id, user_from_id, user_to_id, content, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
		message.Id,
		message.ChannelID,
		message.UserFromID,
		message.UserToID,
		message.Content,
		message.CreatedAt,
	)

	return err
}
