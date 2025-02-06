package domain

// Stat is a struct that represents the remaining resource indicator of the im gateway itself
type Stat struct {
	// The remaining value of the number of the im gateway connect num
	ConnectNum float64
	// The remaining value of the number of the im gateway message bytes sending and receiving per second
	MessageBytes float64
}

// Clone returns a new Stat with the same value as the current Stat
func (s *Stat) Clone() *Stat {
	newStat := &Stat{
		MessageBytes: s.MessageBytes,
		ConnectNum:   s.ConnectNum,
	}
	return newStat
}

// Avg calculates the average value of the current Stat
func (s *Stat) Avg(num float64) {
	s.ConnectNum /= num
	s.MessageBytes /= num
}

// Add adds the value of the input Stat to the current Stat
func (s *Stat) Add(st *Stat) {
	if st == nil {
		return
	}
	s.ConnectNum += st.ConnectNum
	s.MessageBytes += st.MessageBytes
}

// Sub subtracts the value of the input Stat from the current Stat
func (s *Stat) Sub(st *Stat) {
	if st == nil {
		return
	}
	s.ConnectNum -= st.ConnectNum
	s.MessageBytes -= st.MessageBytes
}
