package gateway

import (
	"fmt"
	"net"
	"reflect"
	"runtime"
	"sync/atomic"

	"github.com/0125nia/Mercury/common/config"
)

var ep *ePool    // epoll pool
var tcpNum int32 // max tcp connection number

func initEpoll(listener *net.TCPListener, f func(c *connection, ep *epoller)) {
	ep = newEPool(listener, f)
	// todo
}

// epoll pool
type ePool struct {
	eChan chan *connection

	eSize    int
	done     chan struct{}
	listener *net.TCPListener
	f        func(c *connection, ep *epoller)
}

func newEPool(listener *net.TCPListener, f func(c *connection, ep *epoller)) *ePool {
	return &ePool{
		eChan:    make(chan *connection, config.Config.Gateway.EpollerChanNum),
		eSize:    config.Config.Gateway.EpollerNum,
		done:     make(chan struct{}),
		listener: listener,
		f:        f,
	}
}

// create a gourotine to accept tcp connection
// the number corresponds to the number of cpu cores
func (ep *ePool) createAcceptProcess() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				conn, e := ep.listener.AcceptTCP()
				if !checkTcp() {
					_ = conn.Close()
					continue
				}
				setTCPConifg(conn)
				if e != nil {
					if err, ok := e.(net.Error); ok && err.Timeout() {
						fmt.Printf("accept timeout: %v\n", e)
						continue
					}
					fmt.Printf("accept err: %v\n", e)
				}
				c := connection{
					conn: conn,
					fd:   socketFD(conn),
				}
				ep.addTask(&c)
			}
		}()
	}

}

func (ep *ePool) addTask(c *connection) {
	ep.eChan <- c
}

func getTcpNum() int32 {
	return atomic.LoadInt32(&tcpNum)
}

func checkTcp() bool {
	num := getTcpNum()
	maxTcpNum := config.Config.Gateway.TCPMaxNum
	return num <= maxTcpNum
}

func setTCPConifg(c *net.TCPConn) {
	_ = c.SetKeepAlive(true)
}

// get the file descriptor of the tcp connection
func socketFD(conn *net.TCPConn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(*conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	return int(pfdVal.FieldByName("Sysfd").Int())
}

// epoller is an event poller that encapsulates epoll operations
type epoller struct {
	fd int
}
