package sdk

// connect is a struct that represents a connection between two clients.
type connect struct {
	serverAddr         string
	sendChan, recvChan chan *Message
}

// newConnect creates a new connect struct.
// create a new connect with the given server address and initializes the send and receive channels.
func newConnect(serverAddr string) *connect {
	return &connect{
		serverAddr: serverAddr,
		sendChan:   make(chan *Message),
		recvChan:   make(chan *Message),
	}
}

// send sends a message to the other client.
func (c *connect) send(msg *Message) {
	c.sendChan <- msg
}

// recv receives a message from the other client.
func (c *connect) recv() *Message {
	return <-c.recvChan
}

// close closes the connection.
func (c *connect) close() {
	//todo: close the connection
}
