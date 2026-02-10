package hub

import (
	"strings"
)

type DeleteCommand struct{}

func (c *DeleteCommand) Name() string { return "delete" }

func (c *DeleteCommand) Execute(h *Hub, cmd Command) {
	if cmd.Args == "" {
		h.sendSystem(
			cmd.From,
			"Incorrect number of arguments. Usage: /delete <channel_name>",
		)

		return
	}

	name := strings.TrimSpace(cmd.Args)

	existingChannel := getChannelByName(h.channels, name)
	if existingChannel == nil {
		h.sendSystem(
			cmd.From,
			"Channel with name '"+name+"' does not exist",
		)

		return
	}

	if existingChannel.CreatedBy == nil || existingChannel.CreatedBy.ID != cmd.From.ID {
		h.sendSystem(
			cmd.From,
			"You do not have permissions to delete channel '"+name+"'",
		)

		return
	}

	delete(h.channels, existingChannel.ID)

	h.sendSystem(
		cmd.From,
		"Channel '"+name+"' deleted successfully",
	)
}
