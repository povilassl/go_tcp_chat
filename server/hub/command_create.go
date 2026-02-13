package hub

import (
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
	valid, message := isNameValid(name)
	if !valid {
		h.sendSystemToClient(cmd.From, message)
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
