package domain

import (
	"sync/atomic"
	"unsafe"
)

type Endpoint struct {
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	ActiveScore float64
	StaticScore float64
	Stats       *Stat
	window      *stateWindow
}

func NewEndpoint(ip, port string) *Endpoint {
	ed := &Endpoint{
		Ip:   ip,
		Port: port,
	}
	ed.window = newStateWindow()
	ed.Stats = ed.window.getStat()

	// Start a goroutine to listen to the statChan channel
	go func() {
		// Listen to the statChan channel and append the stat to the window
		for stat := range ed.window.statChan {
			ed.window.appendStat(stat)
			newStat := ed.window.getStat()
			// Safely update the ed.Stats pointer to point to newStat
			// The goal is to ensure data consistency and security in the multithreaded environment
			atomic.SwapPointer((*unsafe.Pointer)((unsafe.Pointer)(ed.Stats)), unsafe.Pointer(newStat))
		}
	}()

	return ed
}

func (ed *Endpoint) CalculateScore(ctx *IpConfContext) {
	// If the result is the same as last time, use the data last time and it isn't updated this time
	if ed.Stats != nil {
		ed.ActiveScore = ed.Stats.CalculateActiveScore()
		ed.StaticScore = ed.Stats.CalculateStaticScore()
	}
}

func (ed *Endpoint) UpdateStat(stat *Stat) {
	ed.window.statChan <- stat
}
