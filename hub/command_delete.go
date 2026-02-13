package hub

import (
	"strings"
)

type DeleteCommand struct{}

func (c *DeleteCommand) Name() string { return "delete" }

func (c *DeleteCommand) Usage() string { return "/delete <channel_name>" }

func (c *DeleteCommand) Execute(h *Hub, cmd Command) {
	if cmd.Args == "" {
		h.sendSystemToClient(
			cmd.From,
			"Incorrect number of arguments. Usage: /delete <channel_name>",
		)

		return
	}

	name := strings.TrimSpace(cmd.Args)

	existingChannel := getChannelByName(h.channels, name)
	if existingChannel == nil {
		h.sendSystemToClient(
			cmd.From,
			"Channel with name '"+name+"' does not exist",
		)

		return
	}

	if existingChannel.CreatedBy == nil || existingChannel.CreatedBy.ID != cmd.From.ID {
		h.sendSystemToClient(
			cmd.From,
			"You do not have permissions to delete channel '"+name+"'",
		)

		return
	}

	delete(h.channels, existingChannel.ID)

	h.sendSystemToClient(
		cmd.From,
		"Channel '"+name+"' deleted successfully",
	)
}
