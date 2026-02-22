package hub

import "fmt"

type Message struct {
	Text        string
	From        *Client
	To          *Client
	ChannelName string
	Type        MessageType
}

type MessageType int

const (
	MessageUnknown MessageType = iota
	MessageChannel
	MessageDirect
	MessageSystem
)

func (m *Message) Format() string {
	var sender string

	switch m.Type {
	case MessageSystem:
		sender = "* System"

	case MessageDirect:
		name := "Unknown"
		if m.From != nil && m.From.DisplayName != "" {
			name = m.From.DisplayName
		}
		sender = fmt.Sprintf("DM from %s", name)

	case MessageChannel:
		name := "Unknown"
		if m.From != nil && m.From.DisplayName != "" {
			name = m.From.DisplayName
		}
		sender = fmt.Sprintf("%s in #%s", name, m.ChannelName)

	default:
		sender = "Unknown"
	}

	return fmt.Sprintf("[%s] %s\r\n", sender, m.Text)
}
