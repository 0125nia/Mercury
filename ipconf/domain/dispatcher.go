package domain

import (
	"sort"
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
	for _, ed := range edList {
		ed.CalculateScore(ctx)
	}

	// sort the endpoint list by active score and static score
	sort.Slice(edList, func(i, j int) bool {
		// compare the active score first
		if edList[i].ActiveScore > edList[j].ActiveScore {
			return true
		}
		// if the active score is equal, then compare the static score
		if edList[i].ActiveScore == edList[j].ActiveScore {
			return edList[i].StaticScore > edList[j].StaticScore
		}
		return false
	})

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
	var ed *Endpoint
	var ok bool
	if ed, ok = dp.candidateTable[event.Key()]; !ok { // ensure the node is not exist
		ed = NewEndpoint(event.Ip, event.Port)
		dp.candidateTable[event.Key()] = ed
	}
	// update the stat
	ed.UpdateStat(&Stat{
		ConnectNum:   event.ConnectNum,
		MessageBytes: event.MessageBytes,
	})
}
