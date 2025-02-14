package tcp

import (
	"bytes"
	"encoding/binary"
)

// DataPgk is a struct that represents a data package.
type DataPgk struct {
	Len  uint32
	Data []byte
}

// Marshal serializes the data package.
func (d *DataPgk) Marshal() []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, d.Len)
	return append(bytesBuffer.Bytes(), d.Data...)
}
