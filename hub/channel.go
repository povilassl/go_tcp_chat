package hub

import (
	"sync/atomic"
)

const limitOfChannelsPerUser = 3

var nextChannelID uint64 = 0

type Channel struct {
	ID        uint64
	Name      string
	Members   map[uint64]*Client
	CreatedBy *Client
}

func NewChannel(name string, createdBy *Client) *Channel {
	id := atomic.AddUint64(&nextChannelID, 1)

	return &Channel{
		ID:        id,
		Name:      name,
		Members:   make(map[uint64]*Client),
		CreatedBy: createdBy,
	}
}

func getChannelByName(channel map[uint64]*Channel, name string) *Channel {
	for _, c := range channel {
		if c.Name == name {
			return c
		}
	}

	return nil
}

// func limitOfChannelsReached(channel map[uint64]*Channel, name string) bool {
// 	count := 0
// 	for _, c := range channel {
// 		if c.CreatedBy != nil && c.CreatedBy.Name == name {
// 			count++
// 		}
// 	}

// 	return count >= limitOfChannelsPerUser
// }
