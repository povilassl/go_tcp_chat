package mysql

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

type ChannelRepository struct {
	db *sqlx.DB
}

func NewChannelRepository(db *sqlx.DB) *ChannelRepository {
	return &ChannelRepository{
		db: db,
	}
}

func (r *ChannelRepository) Create(channel *entity.Channel) error {
	_, err := r.db.Exec("INSERT INTO channels (id, channel_name, created_by_id, created_at) VALUES ($1, $2, $3, $4)",
		channel.ID,
		channel.ChannelName,
		channel.CreatedByID,
		channel.CreatedAt)

	return err
}

func (r *ChannelRepository) GetByName(name string) (*entity.Channel, error) {
	channel := entity.Channel{}
	err := r.db.Get(&channel, "SELECT * FROM channels WHERE channel_name = $1", name)

	if err != nil {
		return nil, err
	}

	return &channel, nil
}

func (r *ChannelRepository) GetByUserID(userID uuid.UUID) (*[]entity.Channel, error) {
	channels := []entity.Channel{}
	err := r.db.Select(&channels, "SELECT * FROM channels WHERE created_by_id = $1", userID)

	if err != nil {
		return nil, err
	}

	return &channels, nil
}

func (r *ChannelRepository) Get(limit int, offset int) (*[]entity.Channel, error) {
	channels := []entity.Channel{}
	err := r.db.Select(&channels, "SELECT * FROM channels LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}

	return &channels, nil
}

func (r *ChannelRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM channels WHERE id = $1", id)

	return err
}
