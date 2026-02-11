package hub

import (
	"regexp"
	"strings"
)

type CreateCommand struct{}

func (c *CreateCommand) Name() string { return "create" }

func (c *CreateCommand) Usage() string { return "/create <channel_name>" }

func (c *CreateCommand) Execute(h *Hub, cmd Command) {
	if cmd.Args == "" {
		h.sendSystemToClient(
			cmd.From,
			"Incorrect number of arguments. Usage: /create <channel_name>",
		)

		return
	}

	name := strings.TrimSpace(cmd.Args)

	if len(name) == 0 || len(name) > 14 {
		h.sendSystemToClient(
			cmd.From,
			"Channel name must be between 1 and 14 characters long",
		)

		return
	}

	if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name) {
		h.sendSystemToClient(
			cmd.From,
			"Channel Name must contain only letters and numbers",
		)

		return
	}

	existingChannel := getChannelByName(h.channels, name)
	if existingChannel != nil {
		h.sendSystemToClient(
			cmd.From,
			"Channel with this name already exists",
		)

		return
	}

	if limitOfChannelsReached(h.channels, cmd.From.Name) {
		h.sendSystemToClient(
			cmd.From,
			"You have reached the limit of channels you can create",
		)

		return
	}

	//TODO maybe index by name
	channel := NewChannel(name, cmd.From)
	h.channels[channel.ID] = channel

	h.sendSystemToClient(
		cmd.From,
		"Channel '"+name+"' created successfully",
	)
}
