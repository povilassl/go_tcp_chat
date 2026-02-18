package hub

import "fmt"

type GetCommand struct{}

func (c *GetCommand) Name() string { return "get" }

func (c *GetCommand) Usage() string { return "/get" }

func (c *GetCommand) BaseErrorMessage() string { return "Error getting channels" }

func (c *GetCommand) Execute(h *Hub, cmd Command) {
	if !h.RequireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	if len(h.channels) == 0 {
		h.sendSystemToClient(
			cmd.From,
			"No channels currently available",
		)

		return
	}

	ret := "Channels:\r\n"
	for _, channel := range h.channels {
		ret += fmt.Sprintf(" - #%s | Members: %d\r\n", channel.Name, len(channel.Members))
	}

	h.sendSystemToClient(
		cmd.From,
		ret,
	)
}
