package hub

import (
	"fmt"
)

func formatMessage(m *Message) string {
	var sender string

	switch m.Type {
	case MessageSystem:
		sender = "System"

	case MessageDirect:
		sender = fmt.Sprintf("DM from %s", m.From.Name)

	case MessagePublic:
		if m.From.Name != "" {
			sender = m.From.Name
		} else {
			sender = "Guest"
		}

	case MessageChannel:
		sender = fmt.Sprintf("%s in Channel", m.From.Name) //TODO

	default:
		sender = "Unknown"
	}

	return fmt.Sprintf("[%s] %s\r\n", sender, m.Text)
}
