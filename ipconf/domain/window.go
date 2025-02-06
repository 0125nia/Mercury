package domain

const (
	windowSize = 5
)

// stateWindow is a sliding window that stores the state of the endpoint
type stateWindow struct {
	stateQueue []*Stat
	statChan   chan *Stat
	sumStat    *Stat
	idx        int64
}

func newStateWindow() *stateWindow {
	return &stateWindow{
		stateQueue: make([]*Stat, windowSize),
		statChan:   make(chan *Stat),
		sumStat:    &Stat{},
	}
}

// getStat returns the average value of the state in the window
func (sw *stateWindow) getStat() *Stat {
	res := sw.sumStat.Clone()
	res.Avg(windowSize)
	return res
}

// appendStat appends the state to the window
func (sw *stateWindow) appendStat(s *Stat) {
	// sub the oldest stat
	sw.sumStat.Sub(sw.stateQueue[sw.idx%windowSize])
	// add the new stat
	sw.stateQueue[sw.idx%windowSize] = s
	// calculate the new state sum
	sw.sumStat.Add(s)
	// move the index
	sw.idx++
}
