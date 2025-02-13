package sdk

import "net"

const (
	MsgTypeText = "text"
)

// Message is a struct that represents a message that is sent between clients.
type Message struct {
	Type    string
	Name    string
	ToId    string
	FromId  string
	Content string
	Session string
}

// Chat is a struct that represents a chat between two clients.
type Chat struct {
	Name      string
	ID        string
	SessionID string
	conn      *connect
}

// NewChat creates a new chat struct.
func NewChat(ip net.IP, port int, name, id, sessionId string) *Chat {
	return &Chat{
		Name:      name,
		ID:        id,
		SessionID: sessionId,
		conn:      newConnect(ip, port),
	}
}

// Send sends a message to the other client.
func (c *Chat) Send(msg *Message) {
	c.conn.send(msg)
}

// Recv receives a message from the other client.
func (c *Chat) Recv() <-chan *Message {
	return c.conn.recv()
}

// Close release resource of the chat.
func (c *Chat) Close() {
	c.conn.close()
}
