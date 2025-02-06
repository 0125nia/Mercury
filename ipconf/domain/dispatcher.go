package domain

import (
	"sync"

	"github.com/0125nia/Mercury/ipconf/source"
)

type Dispatcher struct {
	candidateTable map[string]*Endpoint
	sync.RWMutex
}

var dp *Dispatcher

func Init() {
	dp = &Dispatcher{}
	dp.candidateTable = make(map[string]*Endpoint)
	//
	go func() {
		for event := range source.EventChan() {
			switch event.Type {
			case source.AddNodeEvent:
				// todo add Node
			case source.DelNodeEvent:
				// todo del Node
			}
		}
	}()

}

func Dispatch(ctx *IpConfContext) {

}
