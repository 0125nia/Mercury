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

	go func() { // event-driven
		for event := range source.EventChan() {
			switch event.Type {
			case source.AddNodeEvent:
				dp.addNode(event)
			case source.DelNodeEvent:
				dp.delNode(event)
			}
		}
	}()

}

// Dispatch
func Dispatch(ctx *IpConfContext) []*Endpoint {
	// get candidate endpoint list
	edList := dp.getCandidateList()

	// calculate the score of each endpoint

	// sort the endpoint list

	return edList
}

// getCandidateList Convert the candidateTable to a list
func (d *Dispatcher) getCandidateList() []*Endpoint {
	dp.RLock()
	defer dp.RUnlock()
	res := make([]*Endpoint, 0, len(dp.candidateTable))
	for _, ed := range dp.candidateTable {
		res = append(res, ed)
	}
	return res
}

// delNode is a function to handle the deleted event
func (d *Dispatcher) delNode(event *source.Event) {
	dp.Lock()
	defer dp.Unlock()
	delete(dp.candidateTable, event.Key())
}

// addNode is a function to handle the new event
func (d *Dispatcher) addNode(event *source.Event) {
	dp.Lock()
	defer dp.Unlock()
	// todo add Node
	// var ed *Endpoint
	// var ok bool
	// if ed, ok = dp.candidateTable[event.Key()]; !ok { // ensure the node is not exist

	// }
}
