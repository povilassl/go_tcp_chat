package hub

import (
	"fmt"
	"strings"
)

type NameCommand struct{}

func (c *NameCommand) Name() string { return "name" }

func (c *NameCommand) Usage() string { return "/name <new_name>" }

func (c *NameCommand) BaseErrorMessage() string { return "Error changing name" }

func (c *NameCommand) Execute(h *Hub, cmd Command) {
	if !h.RequireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	args, ok := h.GetArgs(cmd, 1, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	user, err := h.userService.GetByID(cmd.From.UserID)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	originalName := cmd.From.DisplayName
	newName := strings.TrimSpace(args[0])

	err = h.userService.Rename(user, &newName)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	cmd.From.DisplayName = newName

	members, err := h.channelService.GetMembersByUserID(cmd.From.UserID)
	if err != nil {
		return
	}

	msg := Message{
		Text: fmt.Sprintf("%s is now known as %s", originalName, newName),
		Type: MessageSystem,
	}

	h.sendToUserIDs(*members, msg, &cmd.From.UserID)
}
