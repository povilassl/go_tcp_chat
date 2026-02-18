package hub

type Command struct {
	Name string
	Args string
	From *Client
}

type CommandHandler interface {
	Name() string
	Usage() string
	BaseErrorMessage() string
	Execute(h *Hub, cmd Command)
}
