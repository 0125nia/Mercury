package source

import "github.com/0125nia/Mercury/ipconf/discovery"

var eventChan chan *Event

type EventType string

const (
	AddNodeEvent EventType = "addNode"
	DelNodeEvent EventType = "delNode"
)

type Event struct {
	Type EventType
	Ip   string
	Port string
	// todo add more about event here
}

func NewEvent(etype EventType, ed *discovery.Endpoint) *Event {
	if ed == nil {
		return nil
	}

	return &Event{
		Type: etype,
		Ip:   ed.Ip,
		Port: ed.Port,
	}
}
