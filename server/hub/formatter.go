package hub

import (
	"fmt"
)

func formatMessage(m *Message, c *Client) string {
	var sender string

	if m.System {
		sender = "System"
	} else if c.Name != "" {
		sender = m.From.Name
	} else {
		sender = "Guest"
	}

	return fmt.Sprintf("[%s] %s\r\n", sender, m.Text)
}
