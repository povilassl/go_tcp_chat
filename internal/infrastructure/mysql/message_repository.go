package mysql

import (
	"github.com/google/uuid"
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
	_, err := r.db.Exec("INSERT INTO messages (id, channel_id, user_from_id, user_to_id, content, created_at) VALUES (?, ?, ?, ?, ?, ?)",
		message.Id,
		message.ChannelID,
		message.UserFromID,
		message.UserToID,
		message.Content,
		message.CreatedAt,
	)

	return err
}

func (r *MessageRepository) DeleteByChannelID(channelID uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM messages WHERE channel_id = ?", channelID)

	return err
}
