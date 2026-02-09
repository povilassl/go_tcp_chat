package hub

type Message struct {
	Text string
	From *Client
	To   *Client
	Type MessageType
}

type MessageType int

const (
	MessagePublic MessageType = iota
	MessageDirect
	MessageSystem
)
