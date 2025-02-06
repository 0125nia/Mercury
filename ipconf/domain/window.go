package domain

type stateWindow struct {
	stateQueue []*Stat
	statChan   chan *Stat
	sumStat    *Stat
	idx        int64
}
