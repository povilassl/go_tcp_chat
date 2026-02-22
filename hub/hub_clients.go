package hub

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

func (h *Hub) IsClientTracked(c *Client) bool {
	if _, exists := h.pendingClients[c.ID]; exists {
		return true
	}

	if c.UserID != uuid.Nil {
		if existing, exists := h.clientsByUserID[c.UserID]; exists && existing.ID == c.ID {
			return true
		}
	}

	return false
}

func (h *Hub) handleConnect(c *Client) {
	h.pendingClients[c.ID] = c
}

func (h *Hub) handleDisconnect(ev DisconnectEvent) {
	existingClient := ev.Client
	if c, ok := h.pendingClients[ev.Client.ID]; ok {
		existingClient = c
		delete(h.pendingClients, ev.Client.ID)
	}

	if existingClient.UserID != uuid.Nil {
		if c, ok := h.clientsByUserID[existingClient.UserID]; ok && c.ID == existingClient.ID {
			delete(h.clientsByUserID, existingClient.UserID)
		}
	}

	fmt.Printf("Disconnecting client with ID %s. Reason: %s\r\n", existingClient.ID, ev.Reason)

	close(existingClient.Send)
}

func (h *Hub) bindClientUser(client *Client, user *entity.User) error {
	if client == nil || user == nil {
		return fmt.Errorf("Invalid user or client")
	}

	if existing := h.clientsByUserID[user.ID]; existing != nil && existing.ID != client.ID {
		h.handleDisconnect(DisconnectEvent{
			Client: existing,
			Reason: "duplicate login",
		})
	}

	client.UserID = user.ID
	client.DisplayName = user.Nickname
	h.clientsByUserID[user.ID] = client

	delete(h.pendingClients, client.ID)

	return nil
}

func (h *Hub) unbindClientUser(from *Client) error {
	if from.UserID == uuid.Nil {
		return fmt.Errorf("Client is not bound to any user")
	}

	if _, exists := h.clientsByUserID[from.UserID]; !exists {
		return fmt.Errorf("Client is not bound to any user")
	}

	delete(h.clientsByUserID, from.UserID)
	from.UserID = uuid.Nil
	from.DisplayName = ""

	return nil
}

func (h *Hub) findClientByName(s string) *Client {
	for _, c := range h.clientsByUserID {
		if c.DisplayName == s {
			return c
		}
	}

	return nil
}

func (h *Hub) Shutdown() {
	fmt.Println("Disconnecting all clients...")

	h.disconnectClients(h.pendingClients, "server shutdown")
	h.disconnectClients(h.clientsByUserID, "server shutdown")
}

func (h *Hub) disconnectClients(clients map[uuid.UUID]*Client, reason string) {
	for _, c := range clients {
		h.sendSystemToClient(c, "Server is shutting down. Disconnecting all clients.\r\n")

		h.handleDisconnect(DisconnectEvent{
			Client: c,
			Reason: reason,
		})
	}
}
