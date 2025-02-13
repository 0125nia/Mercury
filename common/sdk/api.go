package sdk

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
	Id        string
	SessionId string
	conn      *connect
}

// NewChat creates a new chat struct.
func NewChat(serverAddr, name, id, sessionId string) *Chat {
	return &Chat{
		Name:      name,
		Id:        id,
		SessionId: sessionId,
		conn:      newConnect(serverAddr),
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
