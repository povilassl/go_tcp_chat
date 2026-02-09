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
	}

	return fmt.Sprintf("[%s] %s\r\n", sender, m.Text)
}
