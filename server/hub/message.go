package hub

type Message struct {
	Text    string
	From    *Client
	To      *Client
	Channel *Channel
	Type    MessageType
}

type MessageType int

const (
	MessageUnknown MessageType = iota
	MessagePublic
	MessageChannel
	MessageDirect
	MessageSystem
)

func (msg Message) getRecipients() []*Client {
	if msg.To != nil {
		return []*Client{msg.To}
	}

	if msg.Channel != nil {
		recipients := make([]*Client, 0, len(msg.Channel.Members))
		for _, c := range msg.Channel.Members {
			if msg.From != nil && c.ID == msg.From.ID {
				continue
			}
			recipients = append(recipients, c)
		}

		return recipients
	}

	return nil
}
