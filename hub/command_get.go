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

	channels, err := h.channelService.Get(100, 0)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	if len(*channels) == 0 {
		h.sendSystemToClient(
			cmd.From,
			"No channels currently available",
		)
		return
	}

	ret := "Channels:\r\n"
	for _, channel := range *channels {
		//TODO: add members count
		// ret += fmt.Sprintf(" - #%s | Members: %d\r\n", channel.ChannelName, len(channel.))
		ret += fmt.Sprintf(" - #%s \r\n", channel.ChannelName)
	}

	h.sendSystemToClient(
		cmd.From,
		ret,
	)
}
