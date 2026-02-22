package hub

import (
	"fmt"
	"strings"
)

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
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), "No channels currently available"),
		)
		return
	}

	memberCounts, err := h.channelService.GetMemberCounts()
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	var ret strings.Builder
	ret.WriteString("Channels:\r\n")

	for _, channel := range *channels {
		count := memberCounts[channel.ID]
		ret.WriteString(fmt.Sprintf(" - #%s | Members: %d\r\n", channel.ChannelName, count))
	}

	h.sendSystemToClient(
		cmd.From,
		ret.String(),
	)
}
