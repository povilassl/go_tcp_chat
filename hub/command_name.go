package hub

import (
	"fmt"
	"strings"

	"github.com/povilassl/tcp_chat/hub/helpers"
)

type NameCommand struct{}

func (c *NameCommand) Name() string { return "name" }

func (c *NameCommand) Usage() string { return "/name <new_name>" }

func (c *NameCommand) BaseErrorMessage() string { return "Error changing name" }

func (c *NameCommand) Execute(h *Hub, cmd Command) {

	if cmd.Args == "" {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Incorrect number of arguments. Usage: %s", c.BaseErrorMessage(), c.Usage()),
		)

		return
	}

	originalName := cmd.From.Name
	newName := strings.TrimSpace(cmd.Args)

	valid, message := helpers.IsNicknameValid(newName)
	if !valid {
		h.sendSystemToClient(cmd.From, message)
		return
	}

	cmd.From.Rename(newName)
	h.sendSystemGlobalBroadcast(fmt.Sprintf("%s is now known as %s", originalName, newName))
}
