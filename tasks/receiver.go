package tasks

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

type Receiver struct {
	addr    *net.UDPAddr
	manager *CopyManager
}

func NewReceiver(
	addr *net.UDPAddr,
	manager *CopyManager,
) (*Receiver, error) {
	if addr == nil || !addr.IP.IsMulticast() {
		return nil, errors.New("require multicast address")
	}
	if manager == nil {
		return nil, errors.New("require copy manager")
	}

	return &Receiver{addr, manager}, nil
}

func (r *Receiver) Start(group *sync.WaitGroup) {
	conn, err := net.ListenMulticastUDP("udp", nil, r.addr)
	if err != nil {
		panic(err)
	}

	defer func(conn *net.UDPConn) { _ = conn.Close() }(conn)
	defer group.Done()

	buf := make([]byte, 1024)
	fmt.Println("receiver work")
	for {
		_, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		r.manager.update(src.String())
	}
}
