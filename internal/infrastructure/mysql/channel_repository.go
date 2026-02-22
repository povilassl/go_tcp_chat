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
	_, err := r.db.Exec("INSERT INTO channels (id, channel_name, created_by_id, created_at) VALUES (?, ?, ?, ?)",
		channel.ID,
		channel.ChannelName,
		channel.CreatedByID,
		channel.CreatedAt)

	return err
}

func (r *ChannelRepository) GetByName(name string) (*entity.Channel, error) {
	channel := entity.Channel{}
	err := r.db.Get(&channel, "SELECT * FROM channels WHERE channel_name = ?", name)

	if err != nil {
		return nil, err
	}

	return &channel, nil
}

func (r *ChannelRepository) GetCreatedByUserID(userID uuid.UUID) (*[]entity.Channel, error) {
	channels := []entity.Channel{}
	err := r.db.Select(&channels, "SELECT * FROM channels WHERE created_by_id = ?", userID)

	if err != nil {
		return nil, err
	}

	return &channels, nil
}

func (r *ChannelRepository) Get(limit int, offset int) (*[]entity.Channel, error) {
	channels := []entity.Channel{}
	err := r.db.Select(&channels, "SELECT * FROM channels LIMIT ? OFFSET ?", limit, offset)

	if err != nil {
		return nil, err
	}

	return &channels, nil
}

func (r *ChannelRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM channels WHERE id = ?", id)

	return err
}

func (r *ChannelRepository) RemoveAllMembers(channelID uuid.UUID) error {
	_, err := r.db.Exec(
		"DELETE FROM channel_members WHERE channel_id = ?",
		channelID,
	)

	return err
}

func (r *ChannelRepository) AddMember(userID, channelID uuid.UUID) error {
	_, err := r.db.Exec(
		"INSERT IGNORE INTO channel_members (user_id, channel_id) VALUES (?, ?)",
		userID,
		channelID,
	)

	return err
}

func (r *ChannelRepository) RemoveMember(userID, channelID uuid.UUID) error {
	_, err := r.db.Exec(
		"DELETE FROM channel_members WHERE user_id = ? AND channel_id = ?",
		userID,
		channelID,
	)

	return err
}

func (r *ChannelRepository) GetMembers(channelID uuid.UUID) (*[]uuid.UUID, error) {
	userIDs := []uuid.UUID{}
	err := r.db.Select(&userIDs, "SELECT user_id FROM channel_members WHERE channel_id = ?", channelID)

	if err != nil {
		return nil, err
	}

	return &userIDs, nil
}

func (r *ChannelRepository) GetMembersByUserID(userID uuid.UUID) (*[]uuid.UUID, error) {
	userIDs := []uuid.UUID{}
	err := r.db.Select(&userIDs, `
		SELECT DISTINCT cm2.user_id
		FROM channel_members cm1
		JOIN channel_members cm2 ON cm1.channel_id = cm2.channel_id
		WHERE cm1.user_id = ?
		AND cm2.user_id != ?
	`, userID, userID)

	if err != nil {
		return nil, err
	}

	return &userIDs, nil
}

func (r *ChannelRepository) GetByUserID(userID uuid.UUID) (*[]entity.Channel, error) {
	channels := []entity.Channel{}
	err := r.db.Select(&channels, "SELECT c.* FROM channels c JOIN channel_members cm ON c.id = cm.channel_id WHERE cm.user_id = ?", userID)

	if err != nil {
		return nil, err
	}

	return &channels, nil
}

func (r *ChannelRepository) GetMemberCounts() (map[uuid.UUID]int, error) {
	type countRow struct {
		ChannelID uuid.UUID `db:"channel_id"`
		Count     int       `db:"count"`
	}

	rows := []countRow{}
	err := r.db.Select(&rows, "SELECT channel_id, COUNT(*) as count FROM channel_members GROUP BY channel_id")

	if err != nil {
		return nil, err
	}

	counts := make(map[uuid.UUID]int)
	for _, row := range rows {
		counts[row.ChannelID] = row.Count
	}

	return counts, nil
}
