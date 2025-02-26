package gateway

import (
	"fmt"
	"log"
	"net"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/0125nia/Mercury/common/config"
	"golang.org/x/sys/unix"
)

var epollerPool *ePool // epoll pool
var tcpNum int32       // max tcp connection number

func initEpoll(listener *net.TCPListener, f func(c *connection, ep *epoller)) {
	setLimit()
	epollerPool = newEPool(listener, f)
	epollerPool.createAcceptProcess()
	epollerPool.startEPool()
}

// epoll pool
type ePool struct {
	eChan    chan *connection
	tables   sync.Map
	eSize    int
	done     chan struct{}
	listener *net.TCPListener
	f        func(c *connection, ep *epoller)
}

func newEPool(listener *net.TCPListener, f func(c *connection, ep *epoller)) *ePool {
	return &ePool{
		eChan:    make(chan *connection, config.Config.Gateway.EpollerChanNum),
		eSize:    config.Config.Gateway.EpollerNum,
		tables:   sync.Map{},
		done:     make(chan struct{}),
		listener: listener,
		f:        f,
	}
}

// set the limit of the number of file descriptors
func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	log.Printf("set cur limit: %d", rLimit.Cur)
}

// create a gourotine to accept tcp connection
// the number corresponds to the number of cpu cores
func (ep *ePool) createAcceptProcess() {
	for range runtime.NumCPU() {
		go func() {
			for {
				conn, e := ep.listener.AcceptTCP()
				if !checkTcp() {
					_ = conn.Close()
					continue
				}
				setTCPConifg(conn)
				if e != nil {
					// handle timeout error
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

// start the epoll pool
func (ep *ePool) startEPool() {
	for range ep.eSize {
		go ep.startEProc()
	}
}

// start the epoll process
func (ep *ePool) startEProc() {
	epoller, err := newEpoller()
	if err != nil {
		panic(err)
	}
	// listen the event that connection creating
	go func() {
		for {
			select {
			case <-ep.done:
				return
			case c := <-ep.eChan:
				addTCPNum()
				fmt.Printf("tcpNum: %d\n", tcpNum)
				if err := epoller.add(c); err != nil {
					fmt.Printf("epoll add err: %v\n", err)
					_ = c.conn.Close()
					continue
				}
				fmt.Printf("EpollerPool new connection[%v] tcpSize:%d\n", c.RemoteAddr(), tcpNum)
			}
		}
	}()

	for {
		select {
		case <-ep.done:
			return
		default:
			connections, err := epoller.wait(200) // 200ms per polling
			if err != nil && err != syscall.EINTR {
				fmt.Printf("failed to epoll wait %v\n", err)
				continue
			}
			for _, conn := range connections {
				if conn == nil {
					break
				}
				ep.f(conn, epoller)
			}
		}
	}

}

// add a task to the epoll pool
func (ep *ePool) addTask(c *connection) {
	ep.eChan <- c
}

func addTCPNum() {
	atomic.AddInt32(&tcpNum, 1)
}

func getTcpNum() int32 {
	return atomic.LoadInt32(&tcpNum)
}

func checkTcp() bool {
	num := getTcpNum()
	maxTcpNum := config.Config.Gateway.TCPMaxNum
	return num <= maxTcpNum
}

// set the tcp connection configuration
func setTCPConifg(c *net.TCPConn) {
	// set the keepalive of the tcp connection
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

// newEpoller creates a new epoll instance
func newEpoller() (*epoller, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &epoller{
		fd: fd,
	}, nil
}

// add adds a connection to the epoll instance
func (epoller *epoller) add(c *connection) error {
	fd := c.fd

	err := unix.EpollCtl(epoller.fd, unix.EPOLL_CTL_ADD, fd, &unix.EpollEvent{
		Events: unix.EPOLLIN | unix.EPOLLHUP,
		Fd:     int32(fd),
	})

	if err != nil {
		return err
	}

	epollerPool.tables.Store(fd, c)

	return nil
}

// wait waits for events on the epoll instance
func (epoller *epoller) wait(msec int) ([]*connection, error) {
	events := make([]unix.EpollEvent, config.Config.Gateway.EpollerWaitQueueSize)
	n, err := unix.EpollWait(epoller.fd, events, msec)
	if err != nil {
		return nil, err
	}

	var connections []*connection
	for i := range n {
		if conn, ok := epollerPool.tables.Load(int(events[i].Fd)); ok {
			connections = append(connections, conn.(*connection))
		}
	}
	return connections, nil
}
