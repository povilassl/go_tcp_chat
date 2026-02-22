package hub

import (
	"fmt"
	"slices"
	"strings"
)

type LeaveCommand struct{}

func (c *LeaveCommand) Name() string { return "leave" }

func (c *LeaveCommand) Usage() string { return "/leave <channel_name>" }

func (c *LeaveCommand) BaseErrorMessage() string { return "Error leaving channel" }

func (c *LeaveCommand) Execute(h *Hub, cmd Command) {
	if !h.RequireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	args, ok := h.GetArgs(cmd, 1, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	name := strings.TrimSpace(args[0])

	channel, err := h.channelService.GetByName(name)
	if err != nil || channel == nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Channel '%s' does not exist", c.BaseErrorMessage(), name),
		)
		return
	}

	members, err := h.channelService.GetMembers(channel.ID)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	isMember := slices.Contains(*members, cmd.From.UserID)
	if !isMember {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: You are not a member of channel '%s'", c.BaseErrorMessage(), name),
		)

		return
	}

	if err := h.channelService.RemoveMember(cmd.From.UserID, channel.ID); err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	msg := Message{
		Text: fmt.Sprintf("%s has left the channel #%s", cmd.From.DisplayName, name),
		Type: MessageSystem,
	}

	h.sendToUserIDs(*members, msg, nil)
}
