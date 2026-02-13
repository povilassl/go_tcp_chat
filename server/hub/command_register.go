package hub

import (
	"strings"
)

type RegisterCommand struct{}

func (c *RegisterCommand) Name() string { return "register" }

func (c *RegisterCommand) Usage() string { return "/register <username> <password>" }

func (c *RegisterCommand) Execute(h *Hub, cmd Command) {
	args := strings.SplitN(cmd.Args, " ", 2)

	if len(args) != 2 {
		h.sendSystemToClient(
			cmd.From,
			"Incorrect number of arguments. Usage: /register <username> <password>",
		)

		return
	}

	name := strings.TrimSpace(args[0])
	password := strings.TrimSpace(args[1])

	nameValid, nameValidMessage := isPasswordValid(name)
	if !nameValid {
		h.sendSystemToClient(cmd.From, nameValidMessage)
		return
	}

	passValid, passValidMessage := isNameValid(password)
	if !passValid {
		h.sendSystemToClient(cmd.From, passValidMessage)
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

	channel := NewChannel(name, cmd.From)
	h.channels[channel.ID] = channel

	h.sendSystemToClient(
		cmd.From,
		"Channel '"+name+"' created successfully",
	)
}
