package hub

type Message struct {
	Text   string
	From   *Client
	To     uint64
	System bool
}
