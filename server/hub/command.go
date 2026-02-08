package hub

import "strings"

type Command struct {
	Text string
	From *Client
}

func (h *Hub) ExecuteCommand(cmd Command) {

	parts := strings.SplitN(cmd.Text, " ", 2)

	if strings.HasPrefix(parts[0], "/name") {
		if len(parts) == 2 {
			newName := parts[1]
			cmd.From.Rename(newName)
		} else {
			h.SendDirect(Message{
				Text:   "Too many arguments. Usage: /name <new_name>",
				To:     cmd.From.ID,
				System: true,
			})
		}
	} else if strings.HasPrefix(parts[0], "/quit") {
		h.Disconnect(cmd.From)
	} else {
		h.SendDirect(Message{
			Text:   "Unknown command: " + parts[0],
			To:     cmd.From.ID,
			System: true,
		})
	}
}
