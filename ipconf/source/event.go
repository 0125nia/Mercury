package source

import (
	"fmt"

	"github.com/0125nia/Mercury/common/discovery"
)

var eventChan chan *Event

func EventChan() chan *Event {
	return eventChan
}

type EventType string

const (
	AddNodeEvent EventType = "addNode"
	DelNodeEvent EventType = "delNode"
)

type Event struct {
	Type         EventType
	Ip           string
	Port         string
	ConnectNum   float64
	MessageBytes float64
}

func NewEvent(etype EventType, ed *discovery.EndpointInfo) *Event {
	if ed == nil || ed.MetaData == nil {
		return nil
	}
	var connNum, msgBytes float64
	if data, ok := ed.MetaData["connect_num"]; ok {
		connNum = data.(float64)
	}
	if data, ok := ed.MetaData["message_bytes"]; ok {
		msgBytes = data.(float64)
	}
	return &Event{
		Type:         etype,
		Ip:           ed.Ip,
		Port:         ed.Port,
		ConnectNum:   connNum,
		MessageBytes: msgBytes,
	}
}

func (nd *Event) Key() string {
	return fmt.Sprintf("%s:%s", nd.Ip, nd.Port)
}
