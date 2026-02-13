package hub

import (
	"fmt"
	"strings"
)

type NameCommand struct{}

func (c *NameCommand) Name() string { return "name" }

func (c *NameCommand) Usage() string { return "/name <new_name>" }

func (c *NameCommand) Execute(h *Hub, cmd Command) {

	if cmd.Args == "" {
		h.sendSystemToClient(
			cmd.From,
			"Incorrect number of arguments. Usage: /name <new_name>",
		)

		return
	}

	originalName := cmd.From.Name
	newName := strings.TrimSpace(cmd.Args)

	valid, message := isNameValid(newName)
	if !valid {
		h.sendSystemToClient(cmd.From, message)
		return
	}

	cmd.From.Rename(newName)
	h.sendSystemGlobalBroadcast(fmt.Sprintf("%s is now known as %s", originalName, newName))
}
