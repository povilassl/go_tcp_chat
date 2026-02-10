package hub

import (
	"fmt"
	"regexp"
	"strings"
)

type NameCommand struct{}

func (c *NameCommand) Name() string { return "name" }

func (c *NameCommand) Execute(h *Hub, cmd Command) {

	if cmd.Args == "" {
		h.sendSystem(
			cmd.From,
			"Incorrect number of arguments. Usage: /name <new_name>",
		)

		return
	}

	originalName := cmd.From.Name
	newName := strings.TrimSpace(cmd.Args)

	if len(newName) == 0 || len(newName) > 14 {
		h.sendSystem(
			cmd.From,
			"Name must be between 1 and 14 characters long",
		)

		return
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(newName) {
		h.sendSystem(
			cmd.From,
			"Name must contain only letters and numbers",
		)

		return
	}

	cmd.From.Rename(newName)

	h.handleBroadcast(Message{
		Text: fmt.Sprintf("%s is now known as %s", originalName, newName),
		Type: MessageSystem,
	})
}
