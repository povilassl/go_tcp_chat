package hub

import (
	"fmt"
	"slices"
	"strings"
)

type JoinCommand struct{}

func (c *JoinCommand) Name() string { return "join" }

func (c *JoinCommand) Usage() string { return "/join <channel_name>" }

func (c *JoinCommand) BaseErrorMessage() string { return "Error joining channel" }

func (c *JoinCommand) Execute(h *Hub, cmd Command) {
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
	if isMember {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: You are already a member of channel '%s'", c.BaseErrorMessage(), name),
		)

		return
	}

	if err := h.channelService.AddMember(cmd.From.UserID, channel.ID); err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	h.sendSystemToUserIDs(
		*members,
		fmt.Sprintf("%s has joined the channel #%s", cmd.From.DisplayName, name),
		nil,
	)
}
