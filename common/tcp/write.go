package tcp

import "net"

// SendData sends data to the connection.
func SendData(conn *net.TCPConn, data []byte) error {
	totalLen := len(data)
	writeLen := 0
	for {
		len, err := conn.Write(data[writeLen:])
		if err != nil {
			return err
		}
		writeLen = writeLen + len
		if writeLen >= totalLen {
			break
		}
	}
	return nil
}
