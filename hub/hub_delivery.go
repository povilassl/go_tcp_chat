package hub

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (h *Hub) SendGreeting(client *Client) {
	welcomeMessage := "Welcome to the server!\r\n\r\n"
	welcomeMessage += "-----------------------------\r\n"
	welcomeMessage += "Your client ID is: " + client.ID.String() + "\r\n"
	welcomeMessage += "Time of connection: " + time.Now().Format(time.RFC1123) + "\r\n"
	welcomeMessage += "-----------------------------\r\n"

	h.handleSend(Message{
		Text: welcomeMessage,
		To:   client,
		Type: MessageSystem,
	})
}

func (h *Hub) sendSystemToClient(client *Client, text string) {
	h.handleSend(Message{
		Text: text,
		To:   client,
		Type: MessageSystem,
	})
}

func (h *Hub) disconnectSlowClient(c *Client) {
	fmt.Println("Disconnecting slow client. Client ID:", c.ID.String())
	h.handleDisconnect(DisconnectEvent{
		Client: c,
		Reason: "slow",
	})
}

func (h *Hub) handleSend(msg Message) {
	select {
	case msg.To.Send <- &msg:
	default:
		h.disconnectSlowClient(msg.To)
	}
}

func (h *Hub) sendToUserIDs(userIDs []uuid.UUID, msg Message, excludeUserID *uuid.UUID) {
	for _, userID := range userIDs {
		if excludeUserID != nil && userID == *excludeUserID {
			continue
		}

		c, ok := h.clientsByUserID[userID]
		if !ok {
			continue
		}

		select {
		case c.Send <- &msg:
		default:
			h.disconnectSlowClient(c)
		}
	}
}

func (h *Hub) sendToUserID(userID uuid.UUID, msg Message) {
	h.sendToUserIDs([]uuid.UUID{userID}, msg, nil)
}

func (h *Hub) sendSystemToUserIDs(userIDs []uuid.UUID, text string, excludeUserID *uuid.UUID) {
	h.sendToUserIDs(
		userIDs,
		Message{
			Text: text,
			Type: MessageSystem,
		},
		excludeUserID,
	)
}
