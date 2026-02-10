package hub

type DisconnectEvent struct {
	Client *Client
	Reason string
}
