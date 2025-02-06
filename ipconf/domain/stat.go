package domain

import "math"

// Stat is a struct that represents the remaining resource indicator of the im gateway itself
type Stat struct {
	// The remaining value of the number of the im gateway connect num
	// This belongs to static score
	ConnectNum float64
	// The remaining value of the number of the im gateway message bytes sending and receiving per second
	// This belongs to active score
	MessageBytes float64
}

func (s *Stat) CalculateActiveScore() float64 {
	return getGB(s.MessageBytes)
}

func (s *Stat) CalculateStaticScore() float64 {
	return s.ConnectNum
}

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

// getGB converts the byte value to GB
func getGB(bytes float64) float64 {
	return decimal(bytes / (1 << 30))
}

// decimal rounds the float64 value to two decimal places
func decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}
