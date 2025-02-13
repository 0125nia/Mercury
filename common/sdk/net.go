package sdk

import (
	"fmt"
	"net"
)

// connect is a struct that represents a connection between two clients.
type connect struct {
	conn               *net.TCPConn
	sendChan, recvChan chan *Message
}

// newConnect creates a new connect struct.
// create a new connect with the given server address and initializes the send and receive channels.
func newConnect(ip net.IP, port int) *connect {
	clientConn := &connect{
		sendChan: make(chan *Message),
		recvChan: make(chan *Message),
	}
	addr := &net.TCPAddr{IP: ip, Port: port}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Printf("DialTCP.err=%+v", err)
		return nil
	}
	clientConn.conn = conn
	go func() {
		for {
			// todo read message from the other client
		}
	}()
	return clientConn
}

// send sends a message to the other client.
func (c *connect) send(data *Message) {
	// 直接发送给接收方
	// bytes, _ := json.Marshal(data)
	// todo
}

// recv receives a message from the other client.
func (c *connect) recv() <-chan *Message {
	return c.recvChan
}

// close closes the connection.
func (c *connect) close() {

}
