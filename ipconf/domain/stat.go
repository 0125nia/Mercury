package domain

// Stat is a struct that represents the remaining resource indicator of the im gateway itself
type Stat struct {
	// The remaining value of the number of the im gateway connect num
	ConnectNum float64
	// The remaining value of the number of the im gateway message bytes sending and receiving per second
	MessageBytes float64
}
