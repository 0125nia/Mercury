package domain

type Endpoint struct {
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	ActiveScore float64
	StaticScore float64
	Stats       *Stat
	window      *stateWindow
}

func NewEndpoint(ip, port string) *Endpoint {
	return &Endpoint{
		Ip:   ip,
		Port: port,
	}
}

func (ed *Endpoint) CalculateScore(ctx *IpConfContext) {
	// todo
}

func (ed *Endpoint) UpdateStat(stat *Stat) {
	ed.window.statChan <- stat
}
